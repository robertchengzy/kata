package leetcode

// Merge Sorted Array
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

// First Bad Version
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
