package main

import (
	"fmt"
	"kata/foundatimentals/structure/queue_stack/stack"
)

func main() {
	fmt.Println("good good study, day day up")

	nums := []int{1, 0}
	fmt.Println(stack.FindTargetSumWays2(nums, 1))
}
