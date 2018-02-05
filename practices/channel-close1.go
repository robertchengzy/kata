package main

import (
	"fmt"
	"sync"
)

type T int

func IsClosed(ch <-chan T) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func main() {
	c := make(chan T)
	fmt.Println(IsClosed(c)) // false
	close(c)
	fmt.Println(IsClosed(c)) // true
}

func SafeClose(ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()

	// assume ch != nil here.
	close(ch) // panic if ch is closed
	return true // <=> justClosed = true; return
}

func SafeSend(ch chan T, value T) (closed bool) {
	defer func() {
		if recover() != nil {
			// The return result can be altered
			// in a defer function call.
			closed = true
		}
	}()

	ch <- value // panic if ch is closed
	return false // <=> closed = false; return
}

type MyChannel struct {
	C    chan T
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

type MyChannelx struct {
	C      chan T
	closed bool
	mutex  sync.Mutex
}

func NewMyChannelx() *MyChannelx {
	return &MyChannelx{C: make(chan T)}
}

func (mc *MyChannelx) SafeClosex() {
	mc.mutex.Lock()
	if !mc.closed {
		close(mc.C)
		mc.closed = true
	}
	mc.mutex.Unlock()
}

func (mc *MyChannelx) IsClosedx() bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	return mc.closed
}
