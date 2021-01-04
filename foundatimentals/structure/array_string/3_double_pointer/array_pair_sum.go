package double_pointer

import (
	"sort"
)

/*
数组拆分 I
给定长度为 2n 的数组, 你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从1 到 n 的 min(ai, bi) 总和最大。
示例 1:
输入: [1,4,3,2]
输出: 4
解释: n 等于 2, 最大总和为 4 = min(1, 2) + min(3, 4).
提示:
	1.n 是正整数,范围在 [1, 10000].
	2.数组中的元素范围在 [-10000, 10000].
*/

/*
排序
时间复杂度：O(nlog(n))。排序需要 O(nlog(n)) 的时间。另外会有一次数组的遍历。
空间复杂度：O(1)。仅仅需要常数级的空间.
*/
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

/*
计数排序
时间复杂度：O(n)。需要遍历一次哈希表 arr。
空间复杂度：O(n)。存储一个大小为 n 哈希表 arr 所需要的空间。
*/
func arrayPairSum1(nums []int) int {
	arr := make([]int, 20001)
	lim := 10000
	for _, num := range nums {
		arr[num+lim]++
	}
	d, sum := 0, 0
	for i := -10000; i <= 10000; i++ {
		sum += (arr[i+lim] + 1 - d) / 2 * i
		d = (2 + arr[i+lim] - d) % 2
	}
	return sum
}
