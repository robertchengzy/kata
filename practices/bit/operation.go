package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var x uint8 = 0xAC // x = 10101100
	x = x & 0xF0       // x = 10100000

	for x := 0; x < 100; x++ {
		num := rand.Int()
		if num&1 == 1 {
			fmt.Printf("%d is odd\n", num)
		} else {
			fmt.Printf("%d is even\n", num)
		}
	}

	var a uint8 = 0
	a |= 196
	fmt.Printf("%b\n", a) // prints 11000100

	var a1 uint8 = 0
	a1 |= 196
	a1 |= 210
	fmt.Printf("%d\n", a1) // prints 11000111
}
