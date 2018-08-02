package codewars

/*A Cartesian coordinate system is a coordinate system that specifies each point uniquely in a plane by a pair of
numerical coordinates, which are the signed distances to the point from two fixed perpendicular directed lines,
measured in the same unit of length.

The сoordinates of a point in the grid are written as (x,y). Each point in a coordinate system has eight neighboring
points. Provided that the grid step = 1.

It is necessary to write a function that takes a coordinate on the x-axis and y-axis and returns a list of all the
neighboring points. Points inside list don't have to been sorted (any order is valid).

For Example:

CartesianNeighbor(2,2) -> {{1,1},{1,2},{1,3},{2,1},{2,3},{3,1},{3,2},{3,3}}
CartesianNeighbor(5,7) -> {{6,7},{6,6},{6,8},{4,7},{4,6},{4,8},{5,6},{5,8}}*/

func CartesianNeighbor(x, y int) [][]int {
	data := make([][]int, 0, 8)
	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			if i != x || j != y {
				data = append(data, []int{i, j})
			}
		}
	}
	return data
}

func CartesianNeighbor2(x, y int) [][]int {
	res := make([][]int, 0, 8)
	for xd := -1; xd <= 1; xd++ {
		for yd := -1; yd <= 1; yd++ {
			if xd != 0 || yd != 0 {
				res = append(res, []int{x + xd, y + yd})
			}
		}
	}
	return res
}
