package utils

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

func createNewMiddleware() Middleware {

	// Create a new Middleware
	middleware := func(next http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc which is called by the server eventually
		handler := func(w http.ResponseWriter, r *http.Request) {

			// ... do middleware things

			// Call the next middleware/handler in chain
			next(w, r)
		}

		// Return newly created handler
		return handler
	}

	// Return newly created middleware
	return middleware
}

func test() {
	http.HandleFunc("/", chain(Hello, Logging(), Method("GET")))
	http.ListenAndServe(":8080", nil)
}

func Logging() Middleware {
	return func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			handlerFunc(w, r)
		}
	}
}

func Method(m string) Middleware {
	return func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			handlerFunc(w, r)
		}
	}
}

func chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
