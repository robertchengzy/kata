package codewars

import "math"

// Maximum Subarray 最大子数组和

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

	return maxLeft, maxRight, maxLeft + maxRight
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
