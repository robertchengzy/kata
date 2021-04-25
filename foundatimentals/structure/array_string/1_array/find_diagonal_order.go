package array

/*
对角线遍历
给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。
示例:
	输入:
		[
		 [ 1, 2, 3 ],
		 [ 4, 5, 6 ],
		 [ 7, 8, 9 ]
		]
	输出:[1,2,4,7,5,3,6,8,9]
说明:给定矩阵中的元素总数不会超过 100000 。
*/

func findDiagonalOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}

	n := len(matrix[0])
	arr := make([]int, 0, m*n)
	for l := 0; l < m+n-1; l++ {
		if l%2 == 0 {
			i, j := l, 0
			if l >= m {
				i = m - 1
				j = l - m + 1
			}
			for i >= 0 && j < n {
				arr = append(arr, matrix[i][j])
				i--
				j++
			}
		} else {
			i, j := 0, l
			if l >= n {
				i = l - n + 1
				j = n - 1
			}
			for i < m && j >= 0 {
				arr = append(arr, matrix[i][j])
				i++
				j--
			}
		}
	}

	return arr
}
