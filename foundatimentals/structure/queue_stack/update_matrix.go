package queue_stack

/*
01 矩阵
给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。
两个相邻元素间的距离为 1 。
示例 1:
输入:
0 0 0
0 1 0
0 0 0
输出:
0 0 0
0 1 0
0 0 0
示例 2:
输入:
0 0 0
0 1 0
1 1 1
输出:
0 0 0
0 1 0
1 2 1
注意:
	给定矩阵的元素个数不超过 10000。
	给定矩阵中至少有一个元素是 0。
	矩阵中的元素只在四个方向上相邻: 上、下、左、右。
*/

func updateMatrix(matrix [][]int) [][]int {
	row := len(matrix)
	col := len(matrix[0])
	var queue []int
	type point struct {
		x, y int
	}
	visited := make(map[point]bool)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, i, j)
				p := point{
					i, j,
				}
				visited[p] = true
			}
		}
	}

	var dx = [4]int{-1, 1, 0, 0}
	var dy = [4]int{0, 0, 1, -1}

	for len(queue) > 0 {
		x, y := queue[0], queue[1]
		queue = queue[2:]
		for m := 0; m < 4; m++ {
			tmpX := x + dx[m]
			tmpY := y + dy[m]
			p := point{
				tmpX, tmpY,
			}
			if tmpX >= 0 && tmpY >= 0 && tmpX < row && tmpY < col && !visited[p] {
				matrix[tmpX][tmpY] = matrix[x][y] + 1
				queue = append(queue, tmpX, tmpY)
				visited[p] = true
			}
		}
	}
	return matrix
}
