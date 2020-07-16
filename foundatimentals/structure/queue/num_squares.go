package queue

/*
完全平方数
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
示例 1:
输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.
示例 2:
输入: n = 13
输出: 2
解释: 13 = 4 + 9.
*/

func numSquares(n int) int {
	max := 0
	squareNums := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		squareNums = append(squareNums, i*i)
		max = i
	}

	if max*max == n {
		return 1
	}

	queue := make([]int, 0)
	queue = append(queue, n)

	level := 0
	for len(queue) != 0 {
		level++
		nextQueue := make([]int, 0)
		for _, remainder := range queue {
			for _, square := range squareNums {
				if remainder == square {
					return level
				} else if remainder < square {
					break
				} else {
					nextQueue = append(nextQueue, remainder-square)
				}
			}
		}

		queue = nextQueue
	}

	return level
}
