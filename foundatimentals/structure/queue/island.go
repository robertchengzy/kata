package queue

/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
示例 1:
输入:
[
['1','1','1','1','0'],
['1','1','0','1','0'],
['1','1','0','0','0'],
['0','0','0','0','0']
]
输出: 1
示例 2:
输入:
[
['1','1','0','0','0'],
['1','1','0','0','0'],
['0','0','1','0','0'],
['0','0','0','1','1']
]
输出: 3
解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
*/
type point struct {
	x, y int
}

var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, 1, -1}
var row, col int
var visited map[point]bool

func numIslands(grid [][]byte) int {
	row = len(grid)
	if row == 0 {
		return 0
	}

	col = len(grid[0])
	count := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				BFS(grid, i, j)
				count++
			}
		}
	}
	return count
}

func BFS(grid [][]byte, i, j int) {
	queue := make([]int, 0)
	queue = append(queue, i, j)
	grid[i][j] = '0'
	for len(queue) != 0 {
		i, j := queue[0], queue[1]
		queue = queue[2:]
		for m := 0; m < 4; m++ {
			tmpi := i + dx[m]
			tmpj := j + dy[m]

			p := point{
				x: tmpi,
				y: tmpj,
			}
			if visited[p] {
				continue
			} else {
				visited[p] = true
			}

			if 0 <= tmpi && tmpi < row && 0 <= tmpj && tmpj < col && grid[tmpi][tmpj] == '1' {
				grid[tmpi][tmpj] = '0'
				queue = append(queue, tmpi, tmpj)
			}
		}
	}
}
