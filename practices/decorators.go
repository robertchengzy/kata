package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	var a Adder = adderImpl{}

	fmt.Println(a.Add(1, 2))

	b := AdderFunc(
		func(x, y int) int {
			return x + y
		},
	)

	fmt.Println(Do(b))

	c := WrapCache(&sync.Map{})(Wraplogger(log.New(os.Stdout, "test", 1))(a))
	fmt.Println(c.Add(10, 20))

	var d Adder = AdderFunc(
		func(x, y int) int {
			return x + y
		},
	)

	d = Chain(
		Wraplogger(log.New(os.Stdout, "test", 1)),
		WrapCache(&sync.Map{}),
	)(d)

	fmt.Println(d.Add(10, 20))
}

type Adder interface {
	Add(x, y int) int
}

type adderImpl struct {
}

func (adderImpl) Add(x, y int) int {
	return x + y
}

type AdderFunc func(x, y int) int

func (a AdderFunc) Add(x, y int) int {
	return x + y
}

func Do(adder Adder) int {
	return adder.Add(1, 2)
}

// Middleware function, this function takes in a `Adder` and returns a new `Adder`.
type AdderMiddleware func(Adder) Adder

func Wraplogger(logger *log.Logger) AdderMiddleware {
	return func(a Adder) Adder {
		// Using `AdderFunc` to implement the `Adder` interface.
		fn := func(x, y int) (result int) {
			defer func(t time.Time) {
				logger.Printf("took=%v, x=%v, y=%v, result=%v", time.Since(t), x, y, result)
			}(time.Now())
			// Propogate call to original adder
			return a.Add(x, y)
		}
		// Return a new `Adder` wrapped with the loggin functionality
		return AdderFunc(fn)
	}
}

func WrapCache(cache *sync.Map) AdderMiddleware {
	return func(a Adder) Adder {
		fn := func(x, y int) int {
			key := fmt.Sprintf("x=%dy=%d", x, y)
			val, ok := cache.Load(key)
			if ok {
				return val.(int)
			}

			result := a.Add(x, y)
			cache.Store(key, result)
			return result
		}

		return AdderFunc(fn)
	}
}

func Chain(outer AdderMiddleware, middleware ...AdderMiddleware) AdderMiddleware {
	return func(a Adder) Adder {
		topIndex := len(middleware) - 1
		for i := range middleware {
			a = middleware[topIndex-i](a)
		}
		return outer(a)
	}
}
