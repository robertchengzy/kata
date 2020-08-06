package double_pointer

/*
最大连续1的个数
给定一个二进制数组， 计算其中最大连续1的个数。
示例 1:
输入: [1,1,0,1,1,1]
输出: 3
解释: 开头的两位和最后的三位都是连续1，所以最大连续1的个数是 3.
注意：
	输入的数组只包含 0 和1。
	输入数组的长度是正整数，且不超过 10,000。
*/

func findMaxConsecutiveOnes(nums []int) int {
	l := len(nums)
	if l == 1 && nums[0] == 1 {
		return 1
	}
	left, right := 0, l-1
	sum, leftSum, rightSum := 0, 0, 0
	for left < right {
		if nums[left] == 1 {
			leftSum++
			if leftSum > sum {
				sum = leftSum
			}
		} else {
			if leftSum > sum {
				sum = leftSum
			}
			leftSum = 0
		}
		if nums[right] == 1 {
			rightSum++
			if rightSum > sum {
				sum = rightSum
			}
		} else {
			if rightSum > sum {
				sum = rightSum
			}
			rightSum = 0
		}

		if left+1 == right && nums[left] == 1 && nums[right] == 1 {
			if sum < leftSum+rightSum {
				sum = leftSum + rightSum
			}
		}
		right--
		left++
		if left == right && nums[left] == 1 && nums[right] == 1 {
			if sum < leftSum+rightSum+1 {
				sum = leftSum + rightSum + 1
			}
		}
	}
	return sum
}

func findMaxConsecutiveOnes1(nums []int) int {
	sum, tmp := 0, 0
	for _, num := range nums {
		if num == 1 {
			tmp += 1
		} else {
			tmp = 0
		}
		if sum < tmp {
			sum = tmp
		}
	}

	return sum
}
