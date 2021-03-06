package array

/*
搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
你可以假设数组中无重复元素。
示例 1:
	输入: [1,3,5,6], 5
	输出: 2
示例 2:
	输入: [1,3,5,6], 2
	输出: 1
示例 3:
	输入: [1,3,5,6], 7
	输出: 4
示例 4:
	输入: [1,3,5,6], 0
	输出: 0
*/

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		} else if nums[i] > target {
			return i
		} else if i == len(nums)-1 && nums[i] < target {
			return i + 1
		}
	}
	return 0
}

/*
时间复杂度：O(logn)，其中 n 为数组的长度。二分查找所需的时间复杂度为 O(logn)。
空间复杂度：O(1)。我们只需要常数空间存放若干变量。
*/
func searchInsert1(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := left + (right-left)/2
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
