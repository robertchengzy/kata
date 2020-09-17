package array_string

/*
杨辉三角
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
在杨辉三角中，每个数是它左上方和右上方的数的和。
示例:
	输入: 5
	输出:
	[
		 [1],
		[1,1],
	   [1,2,1],
	  [1,3,3,1],
	 [1,4,6,4,1]
	]
*/

func generate1(numRows int) [][]int {
	ans := make([][]int, numRows)
	if numRows == 0 {
		return ans
	}
	ans[0] = []int{1}
	for i := 1; i < numRows; i++ {
		prevRow := ans[i-1]
		var row []int
		row = append(row, 1)
		for j := 1; j < i; j++ {
			row = append(row, prevRow[j-1]+prevRow[j])
		}
		row = append(row, 1)
		ans[i] = row
	}
	return ans
}
