package codewars

import "math"

// Maximum Subarray 最大子数组和
// a := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}

// 暴力求解法
func FindMaxSubArray(a []int) (int, int, int) {
	length := len(a)
	sum := math.MinInt64
	maxLeft, maxRight := 0, 0
	for i := 0; i < length; i++ {
		curSum := 0
		for j := i; j < length; j++ {
			curSum += a[j]
			if curSum > sum {
				sum = curSum
				maxLeft = i
				maxRight = j
			}
		}
	}

	return maxLeft, maxRight, sum
}

// 分治法
func findMaxCrossingSubArray(a []int, low, mid, high int) (int, int, int) {
	leftSum := math.MinInt64
	leftTemp := 0
	maxLeft := 0
	for i := mid; i >= low; i-- {
		leftTemp += a[i]
		if leftTemp > leftSum {
			leftSum = leftTemp
			maxLeft = i
		}
	}

	rightSum := math.MinInt64
	rightTemp := 0
	maxRight := 0
	for j := mid + 1; j <= high; j++ {
		rightTemp += a[j]
		if rightTemp > rightSum {
			rightSum = rightTemp
			maxRight = j
		}
	}

	return maxLeft, maxRight, leftSum + rightSum
}

func FindMaximumSubArray(a []int, low, high int) (int, int, int) {
	if low == high {
		return low, high, a[low]
	} else {
		mid := (low + high) / 2
		leftLow, leftHigh, leftSum := FindMaximumSubArray(a, low, mid)
		rightLow, rightHigh, rightSum := FindMaximumSubArray(a, mid+1, high)
		crossLow, crossHigh, crossSum := findMaxCrossingSubArray(a, low, mid, high)
		if leftSum > rightSum && leftSum > crossSum {
			return leftLow, leftHigh, leftSum
		} else if rightSum > leftSum && rightSum > crossSum {
			return rightLow, rightHigh, rightSum
		} else {
			return crossLow, crossHigh, crossSum
		}
	}
}

/*
	线性时间算法
    最优方法，时间复杂度O（n）
	和最大的子序列的第一个元素肯定是正数
	因为元素有正有负，因此子序列的最大和一定大于0
*/
func MaxSubArray(a []int) (int, int, int) {
	maxSum := math.MinInt64
	curSum := 0
	left, right := 0, 0

	length := len(a)
	for i := 0; i < length; i++ {
		curSum += a[i]
		if curSum > maxSum {
			right = i
			maxSum = curSum
		}

		if curSum < 0 {
			left = i + 1
			curSum = 0
		}
	}

	return left, right, maxSum
}

func maxSubArray1(a []int) int {
	ans := a[0]
	sum := 0
	for _, num := range a {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}

		ans = int(math.Max(float64(ans), float64(sum)))
	}

	return ans
}
