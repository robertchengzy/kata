package array

import "sort"

/*
合并区间
给出一个区间的集合，请合并所有重叠的区间。
示例 1:
	输入: [[1,3],[2,6],[8,10],[15,18]]
	输出: [[1,6],[8,10],[15,18]]
	解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:
	输入: [[1,4],[4,5]]
	输出: [[1,5]]
	解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/

/*
时间复杂度：O(nlogn)，其中 n 为区间的数量。除去排序的开销，我们只需要一次线性扫描，所以主要的时间开销是排序的 O(nlogn)。
空间复杂度：O(logn)，其中 n 为区间的数量。这里计算的是存储答案之外，使用的额外空间。O(logn) 即为排序所需要的空间复杂度。
*/
func merge(intervals [][]int) [][]int {
	n := len(intervals)
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	for i := 0; i < n-1; i++ {
		if intervals[i][1] >= intervals[i+1][0] {
			if intervals[i+1][1] > intervals[i][1] {
				intervals[i][1] = intervals[i+1][1]
			}
			// 向前合并
			// [1,6], [2, 6] --> [1,6]
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			i--
			n--
		}
	}
	return intervals
}
