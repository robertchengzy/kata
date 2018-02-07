package main

import "fmt"

type MyError string

func (e *MyError) Error() string {
	return string(*e)
}

var ErrBad = MyError("ErrBad")

func bad() bool {
	return false
}

func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p // Will always return a non-nil error.
}

// 解决办法
/*func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
		return p
	}
	return nil
}*/

func main() {
	err := returnsError()
	if err != nil {
		fmt.Println("return non-nil error")
		return
	}
	fmt.Println("return nil")
}