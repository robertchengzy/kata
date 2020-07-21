package stack

/*
目标和
给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。
返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
示例：
输入：nums: [1, 1, 1, 1, 1], S: 3
输出：5
解释：
-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3
一共有5种方法让最终目标和为3。
提示：
	数组非空，且长度不会超过 20 。
	初始的数组的和不会超过 1000 。
	保证返回的最终结果能被 32 位整数存下。
*/

func findTargetSumWays(nums []int, S int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	count := 0
	if length == 1 && S == 0 && nums[0] == 0 {
		return 2
	}

	if length == 1 && (S-nums[0] == 0 || S+nums[0] == 0) {
		return 1
	}

	if length > 1 {
		count += findTargetSumWays(nums[1:], S-nums[0])
		count += findTargetSumWays(nums[1:], S+nums[0])
	}

	return count
}

var count int

func calculate(nums []int, i, sum, S int) {
	if len(nums) == i {
		if sum == S {
			count++
		}
	} else {
		calculate(nums, i+1, sum+nums[i], S)
		calculate(nums, i+1, sum-nums[i], S)
	}
}

func findTargetSumWays2(nums []int, S int) int {
	calculate(nums, 0, 0, S)
	return count
}
