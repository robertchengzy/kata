package main

import (
	// "strconv"
	"sync"
	"fmt"
	//"time"
	"math/rand"
)

func main() {
	m := make(map[int]int)
	go func() {
		for {
			resultLock.RLock()
			_ = m[1]
			resultLock.RUnlock()
		}
	}()
	go func() {
		for {
			resultLock.Lock()
			m[2] = 2
			resultLock.Unlock()
		}
	}()
	select {}
}

var data = make(map[string]string, 0)
var resultLock sync.RWMutex

func addValue(key, value string, resultChan chan bool) {
	resultLock.Lock()
	data[key] = value
	resultLock.Unlock()
	resultChan <- true
}

func getValue(key string, resultChan chan bool) string {
	str := data[key]
	fmt.Println(str)
	resultChan <- true
	return str
}

func randValue(r *rand.Rand) interface{} {
	b := make([]byte, r.Intn(4))
	for i := range b {
		b[i] = 'a' + byte(rand.Intn(26))
	}
	return string(b)
}
