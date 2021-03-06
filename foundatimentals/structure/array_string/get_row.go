package array_string

/*
杨辉三角 II
给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
在杨辉三角中，每个数是它左上方和右上方的数的和。
示例:
	输入: 3
	输出: [1,3,3,1]
进阶：
	你可以优化你的算法到 O(k) 空间复杂度吗？
*/

func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		ans[i] = 1
		for j := i - 1; j > 0; j-- {
			ans[j] = ans[j-1] + ans[j]
		}
	}
	return ans
}
