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

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		arr := make([]int, i+1)
		arr[0] = 1
		if i > 0 {
			arr[i] = 1
		}
		for j := 1; j < i && i > 1; j++ {
			arr[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		ans[i] = arr
	}
	return ans
}
