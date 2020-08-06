package double_pointer

import (
	"math"
	"sort"
)

/*
长度最小的子数组
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
示例：
	输入：s = 7, nums = [2,3,1,2,4,3]
	输出：2
	解释：子数组 [4,3] 是该条件下的长度最小的子数组。
进阶：
如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
*/

/*
时间复杂度：O(n^2)，其中 n 是数组的长度。需要遍历每个下标作为子数组的开始下标，对于每个开始下标，需要遍历其后面的下标得到长度最小的子数组。
空间复杂度：O(1)。

*/
func minSubArrayLen(s int, nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	min := math.MaxInt32
	for i := 0; i < l; i++ {
		sum := 0
		for j := i; j < l; j++ {
			sum += nums[j]
			if sum >= s {
				if min > j-i {
					min = j - i + 1
				}
				break
			}
		}
	}

	if min == math.MaxInt32 {
		return 0
	}
	return min
}

/*
时间复杂度：O(nlogn)，其中 nn 是数组的长度。需要遍历每个下标作为子数组的开始下标，
遍历的时间复杂度是 O(n)，对于每个开始下标，需要通过二分查找得到长度最小的子数组，二分查找得时间复杂度是 O(\log n)O(logn)，
因此总时间复杂度是 O(nlogn)。
空间复杂度：O(n)，其中 nn 是数组的长度。额外创建数组 sums 存储前缀和。
*/
func minSubArrayLen1(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	sums := make([]int, n+1)
	// 为了方便计算，令 size = n + 1
	// sums[0] = 0 意味着前 0 个元素的前缀和为 0
	// sums[1] = A[0] 前 1 个元素的前缀和为 A[0]
	// 以此类推
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	for i := 1; i <= n; i++ {
		target := s + sums[i-1]
		bound := sort.SearchInts(sums, target)
		if bound < 0 {
			bound = -bound - 1
		}
		if bound <= n {
			ans = min(ans, bound-(i-1))
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

/*
时间复杂度：O(n)，其中 n 是数组的长度。指针 start 和 end 最多各移动 n 次。
空间复杂度：O(1)。
*/
func minSubArrayLen2(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= s {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
