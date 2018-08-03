package leetcode

import "math"

// Climbing Stairs 爬楼梯
/*
	假设你正在爬楼梯。需要 n 步你才能到达楼顶。

	每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

	注意：给定 n 是一个正整数。
*/
// Approach 1: Brute Force
// Time complexity : O(2^n). Size of recursion tree will be 2^n
// Space complexity : O(n). The depth of the recursion tree can go upto n.
func climbStairs(n int) int {
	return climbStairsHelp(0, n)
}

func climbStairsHelp(i, n int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	return climbStairsHelp(i+1, n) + climbStairsHelp(i+2, n)
}

// Approach 2: Recursion with memoization
// Time complexity : O(n). Size of recursion tree can go upto n.
// Space complexity : O(n). The depth of recursion tree can go upto n.
func climbStairs2(n int) int {
	memo := make([]int, n+1)
	return climbStairsHelp2(0, n, memo)
}

func climbStairsHelp2(i, n int, memo []int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	if memo[i] > 0 {
		return memo[i]
	}
	memo[i] = climbStairsHelp2(i+1, n, memo) + climbStairsHelp2(i+2, n, memo)
	return memo[i]
}

// Approach 3: Dynamic Programming
// Time complexity : O(n). Single loop upto nn.
// Space complexity : O(n). dpdp array of size nn is used.
func climbStairs3(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// Approach 4: Fibonacci Number
// Time complexity : O(n). Single loop upto nn is required to calculate n^th fibonacci number.
// Space complexity : O(1). Constant space is used.
func climbStairs4(n int) int {
	if n == 1 {
		return 1
	}
	first, second := 1, 2
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}

// Approach 5: Binets Method
// Time complexity : O(log(n)). Traversing on log(n) bits.
// Space complexity : O(1). Constant space is used.
func climbStairs5(n int) int {
	q := [][]int{{1, 1}, {1, 0}}
	res := pow(q, n)
	return res[0][0]
}

func pow(a [][]int, n int) [][]int {
	ret := [][]int{{1, 0}, {0, 1}}
	for n > 0 {
		if n&1 == 1 {
			ret = multiply(ret, a)
		}
		n >>= 1
		a = multiply(a, a)
	}
	return ret
}

func multiply(a, b [][]int) [][]int {
	c := [][]int{{0, 0}, {0, 0}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
		}
	}
	return c
}

// Approach 6: Fibonacci Formula
// Time complexity : O(log(n))O(log(n)). powpow method takes log(n)log(n) time.
// Space complexity : O(1)O(1). Constant space is used.
func climbStairs6(n int) int {
	sqrt5 := math.Sqrt(5)
	fibn := math.Pow((1+sqrt5)/2, float64(n+1)) - math.Pow((1-sqrt5)/2, float64(n+1))
	return (int)(fibn / sqrt5)
}

// Best Time to Buy and Sell Stock 买卖股票的最佳时机
/*
	给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
	如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。
	注意你不能在买入股票前卖出股票。
	示例 1:
	输入: [7,1,5,3,6,4]
	输出: 5
	解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
		 注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
	示例 2:
	输入: [7,6,4,3,1]
	输出: 0
	解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
*/
// Approach 1: Brute Force
// Time complexity : O(n^2). Loop runs n(n-1)/2 times.
// Space complexity : O(1). Only two variables - maxprofit and profit are used.
func maxProfitValue(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

// Approach 2: One Pass
// Time complexity : O(n). Only a single pass is needed.
// Space complexity : O(1). Only two variables are used.
func maxProfitValue2(prices []int) int {
	minPrice := math.MaxInt64
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

// Key observation: prices[2] - prices[0] = prices[2] - prices[1] + prices[1] - prices[0]
func maxProfitValue3(prices []int) int {
	tmp := 0
	max := 0
	for i := 1; i < len(prices); i++ {
		tmp += prices[i] - prices[i-1]
		if tmp < 0 {
			tmp = 0
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}

// Maximum Subarray 最大子序和
/*
	给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

	示例:

	输入: [-2,1,-3,4,-1,2,1,-5,4],
	输出: 6
	解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
	进阶:

	如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/
// This algorithm runs in O(N) time and uses O(1) space.
func maxSubArray(nums []int) int {
	l := len(nums)
	dp := make([]int, l) // dp[i] means the maximum subarray ending with A[i];
	dp[0] = nums[0]
	maxSum := dp[0]
	for i := 1; i < l; i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		maxSum = max(maxSum, dp[i])
	}

	return maxSum
}

func maxSubArray1(nums []int) int {
	maxSum := math.MinInt64
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func maxSubArray2(nums []int) int {
	maxSum := math.MinInt64
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func maxSubArray3(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	ans, max := 0, nums[0]
	for _, v := range nums {
		if ans >= 0 {
			ans += v
		} else {
			ans = v
		}
		if ans > max {
			max = ans
		}
	}
	return max
}

// Maximum Subarray 打家劫舍
/*
	你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
	给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
	示例 1:
	输入: [1,2,3,1]
	输出: 4
	解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
		 偷窃到的最高金额 = 1 + 3 = 4 。
	示例 2:
	输入: [2,7,9,3,1]
	输出: 12
	解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
		 偷窃到的最高金额 = 2 + 9 + 1 = 12 。
*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) < 2 {
		return nums[0]
	}
	pprev := nums[0]
	prev := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		curr := max(pprev+nums[i], prev)
		pprev = prev
		prev = curr
	}

	return max(pprev, prev)
}

func rob2(nums []int) int {
	yes, no := 0, 0
	for _, v := range nums {
		yes, no = no+v, max(yes, no)
	}
	return max(yes, no)
}
