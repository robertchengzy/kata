package array_string

/*
寻找旋转排序数组中的最小值
假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
请找出其中最小的元素。
你可以假设数组中不存在重复元素。
示例 1:
	输入: [3,4,5,1,2]
	输出: 1
示例 2:
	输入: [4,5,6,7,0,1,2]
	输出: 0
*/

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == nums[right] {
			right = right - 1
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

/*
算法
	1.找到数组的中间元素 mid。
	2.如果中间元素 > 数组第一个元素，我们需要在 mid 右边搜索变化点。
	3.如果中间元素 < 数组第一个元素，我们需要在 mid 左边搜索变化点。
	4.当我们找到变化点时停止搜索，当以下条件满足任意一个即可：
		nums[mid] > nums[mid + 1]，因此 mid+1 是最小值。
		nums[mid - 1] > nums[mid]，因此 mid 是最小值。
时间复杂度：和二分搜索一样 O(logN)
空间复杂度：O(1)
*/
func findMin1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left, right := 0, len(nums)
	if nums[0] < nums[right] {
		return nums[0]
	}

	for right >= left {
		mid := left + (left+right)/2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		if nums[mid] > nums[0] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
