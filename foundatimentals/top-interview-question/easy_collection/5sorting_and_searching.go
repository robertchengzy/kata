package leetcode

// Merge Sorted Array 合并两个有序数组
/*
	给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
	说明:
	初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
	你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
	示例:
	输入:
		nums1 = [1,2,3,0,0,0], m = 3
		nums2 = [2,5,6],       n = 3
	输出: [1,2,2,3,5,6]
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	idx := len(nums1) - 1
	i1 := m - 1
	i2 := n - 1
	for i1 >= 0 && i2 >= 0 {
		n1 := nums1[i1]
		n2 := nums2[i2]
		if n2 > n1 {
			nums1[idx] = n2
			i2--
		} else {
			nums1[idx] = n1
			i1--
		}
		idx--
	}

	for i2 >= 0 {
		nums1[idx] = nums2[idx]
		idx--
		i2--
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	current := len(nums1) - 1
	i := m - 1
	j := n - 1
	for i >= 0 || j >= 0 {
		if j < 0 || (i >= 0 && nums1[i] >= nums2[j]) {
			nums1[current] = nums1[i]
			i--
		} else {
			nums1[current] = nums2[j]
			j--
		}
		current--
	}
}

// First Bad Version 第一个错误的版本
/*
	你是产品经理，目前正在带领一个团队开发新的产品。不幸的是，你的产品的最新版本没有通过质量检测。由于每个版本都是基于之前的版本开发的，所以错误的版本之后的所有版本都是错的。
	假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
	你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。
	示例:
		给定 n = 5，并且 version = 4 是第一个错误的版本。
		调用 isBadVersion(3) -> false
		调用 isBadVersion(5) -> true
		调用 isBadVersion(4) -> true
		所以，4 是第一个错误的版本。
*/
// Forward declaration of isBadVersion API.
func isBadVersion(version int) bool {
	return false
}

// Time complexity : O(n). Assume that isBadVersion(version) takes constant time to check if a version is bad. It takes at most n−1 checks, therefore the overall time complexity is O(n).
// Space complexity : O(1).
func firstBadVersion(n int) int {
	for i := 1; i < n; i++ {
		if isBadVersion(i) {
			return i
		}
	}
	return n
}

// Time complexity : O(logn). The search space is halved each time, so the time complexity is O(logn).
// Space complexity : O(1).
func firstBadVersion2(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (left+right)/2 // avoid overflow
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
