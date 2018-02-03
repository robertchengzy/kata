package Foundatimentals

import "math"

/*Integers: Recreation One

Divisors of 42 are : 1, 2, 3, 6, 7, 14, 21, 42. These divisors squared are: 1, 4, 9, 36, 49, 196, 441, 1764.
The sum of the squared divisors is 2500 which is 50 * 50, a square!

Given two integers m, n (1 <= m <= n) we want to find all integers between m and n whose sum of squared divisors
is itself a square. 42 is such a number.

The result will be an array of arrays or of tuples (in C an array of Pair) or a string, each subarray having two
elements, first the number whose squared divisors is a square and then the sum of the squared divisors.

#Examples:

list_squared(1, 250) --> [[1, 1], [42, 2500], [246, 84100]]
list_squared(42, 250) --> [[42, 2500], [246, 84100]]*/

func ListSquared(m, n int) [][]int {
	data := make([][]int, 0)
	for i := m; i <= n; i++ {
		divisorSum := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				divisorSum += j * j
			}
		}
		x := math.Sqrt(float64(divisorSum))
		if math.Trunc(x) == x {
			data = append(data, []int{i, divisorSum})
		}
	}

	return data
}

func ListSquared2(m, n int) [][]int {
	res := make([][]int, 0)
	for i := m; i <= n; i++ {
		s := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				s += j * j
			}
		}
		if math.Mod(math.Sqrt(float64(s)), 1.) == 0 {
			res = append(res, []int{i, s})
		}
	}
	return res
}
