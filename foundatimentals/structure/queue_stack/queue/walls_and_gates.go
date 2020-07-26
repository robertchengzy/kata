package queue

/*
墙与门
题目描述:
你被给定一个 m × n 的二维网格，网格中有以下三种可能的初始化值：
-1 表示墙或是障碍物
0 表示一扇门
INF 无限表示一个空的房间。然后，我们用 231 - 1 = 2147483647 代表 INF。你可以认为通往门的距离总是小于 2147483647 的。
你要给每个空房间位上填上该房间到 最近 门的距离，如果无法到达门，则填 INF 即可。
示例：
给定二维网格：
INF  -1  0  INF
INF INF INF  -1
INF  -1 INF  -1
  0  -1 INF INF
运行完你的函数后，该网格应该变成：
  3  -1   0   1
  2   2   1  -1
  1  -1   2  -1
  0  -1   3   4
*/

var dxx = [4]int{-1, 1, 0, 0}
var dyy = [4]int{0, 0, 1, -1}

func wallsAndGates(rooms [][]int) {
	row := len(rooms)
	if row == 0 {
		return
	}
	queue := make([]int, 0)
	col := len(rooms[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if rooms[i][j] == 0 {
				queue = append(queue, i, j)
			}
		}
	}

	for len(queue) > 0 {
		i, j := queue[0], queue[1]
		queue = queue[2:]
		for m := 0; m < 4; m++ {
			x := i + dxx[m]
			y := j + dyy[m]
			if x >= 0 && y >= 0 && x < row && y < col && rooms[x][y] >= rooms[i][j]+1 {
				rooms[x][y] = rooms[i][j] + 1
				queue = append(queue, x, y)
			}
		}
	}
}

func wallsAndGates2(rooms [][]int) {
	row = len(rooms)
	if row == 0 {
		return
	}
	col = len(rooms[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if rooms[i][j] == 0 {
				dfs(rooms, i, j, 0)
			}
		}
	}
}

func dfs(rooms [][]int, i, j, val int) {
	if i >= 0 && j >= 0 && i < row && j < col && rooms[i][j] >= val {
		rooms[i][j] = val
		dfs(rooms, i+1, j, val+1)
		dfs(rooms, i-1, j, val+1)
		dfs(rooms, i, j+1, val+1)
		dfs(rooms, i, j-1, val+1)
	}
}
