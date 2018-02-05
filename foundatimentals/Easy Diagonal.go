package foundatimentals

import "math/big"

/*In the drawing below we have a part of the Pascal's triangle, lines are numbered from zero (top). The left diagonal in pale blue with only numbers equal to 1 is diagonal zero, then in dark green (1, 2, 3, 4, 5, 6, 7) is diagonal 1, then in pale green (1, 3, 6, 10, 15, 21) is diagonal 2 and so on.

We want to calculate the sum of the binomial coefficients on a given diagonal. The sum on diagonal 0 is 8 (we'll write it S(7, 0), 7 is the number of the line where we start, 0 is the number of the diagonal). In the same way S(7, 1) is 28, S(7, 2) is 56.

Can you write a program which calculate S(n, p) where n is the line where we start and p is the number of the diagonal?

The function will take n and p (with: n >= p >= 0) as parameters and will return the sum.

##Examples:

diagonal(20, 3) => 5985
diagonal(20, 4) => 20349
##Hint: When following a diagonal from top to bottom have a look at the numbers on the diagonal at its right.

##Ref: http://mathworld.wolfram.com/BinomialCoefficient.html*/

func Diagonal(n, p int) int {
	var dis [][]int
	var sum int
	nums := make([]int, 0)
	for i := 0; i < n+1; i++ {
		var linenums []int
		for j := 0; j < i+1; j++ {
			var length = len(nums)
			var value int

			if j == 0 || j == i {
				value = 1
			} else {
				value = nums[length-i] + nums[length-i-1]
			}
			nums = append(nums, value)
			linenums = append(linenums, value)

			if j == p {
				sum += value
			}
		}
		dis = append(dis, linenums)
	}

	return sum
}

func chooseY(n, k int) int {
	var kk = k
	if k > n-k {
		kk = n - k
	}
	var result = 1
	for i := 0; i < kk; i++ {
		result *= n - i
		result /= i + 1
	}
	return result
}

func Diagonal1(n, p int) int {
	return chooseY(n+1, p+1)
}

func Diagonal2(n, p int) int {
	var z big.Int
	return int(z.Binomial(int64(n+1), int64(p+1)).Int64())
}

func Diagonal3(n, p int) int {
	var sum = 1
	for i := 0; i <= p; i++ {
		sum = sum * (n + 1 - i) / (i + 1)
	}
	return sum
}
