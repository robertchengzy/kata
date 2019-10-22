package main

import (
	"fmt"
	"kata/foundatimentals/top-interview-question/codewars"
)

func main() {
	fmt.Println("good good study, day day up")

	a := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	fmt.Println(codewars.FindMaxSubArray(a))
	fmt.Println(codewars.FindMaximumSubArray(a, 0, len(a)-1))
}
