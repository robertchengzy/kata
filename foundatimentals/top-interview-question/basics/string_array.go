package basics

/*
寻找数组的中心索引
给定一个整数类型的数组 nums，请编写一个能够返回数组“中心索引”的方法。

我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。

如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。

示例 1:

输入:
nums = [1, 7, 3, 6, 5, 6]
输出: 3
解释:
索引3 (nums[3] = 6) 的左侧数之和(1 + 7 + 3 = 11)，与右侧数之和(5 + 6 = 11)相等。
同时, 3 也是第一个符合要求的中心索引。
示例 2:

输入:
nums = [1, 2, 3]
输出: -1
解释:
数组中不存在满足此条件的中心索引。
说明:

nums 的长度范围为 [0, 10000]。
任何一个 nums[i] 将会是一个范围在 [-1000, 1000]的整数。
*/
func pivotIndex(nums []int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		left, right := 0, 0
		for m := 0; m < i; m++ {
			left += nums[m]
		}
		for n := i + 1; n < length; n++ {
			right += nums[n]
		}

		if left == right {
			return i
		}
	}

	return -1
}

/*
至少是其他数字两倍的最大数
在一个给定的数组nums中，总是存在一个最大元素 。

查找数组中的最大元素是否至少是数组中每个其他数字的两倍。

如果是，则返回最大元素的索引，否则返回-1。

示例 1:

输入: nums = [3, 6, 1, 0]
输出: 1
解释: 6是最大的整数, 对于数组中的其他整数,
6大于数组中其他元素的两倍。6的索引是1, 所以我们返回1.


示例 2:

输入: nums = [1, 2, 3, 4]
输出: -1
解释: 4没有超过3的两倍大, 所以我们返回 -1.


提示:

nums 的长度范围在[1, 50].
每个 nums[i] 的整数范围在 [0, 100].
*/
func DominantIndex(nums []int) int {
	// 查找最大数
	max := 0
	index := 0
	for i, num := range nums {
		if num > max {
			max = num
			index = i
		}
	}

	var isExist = true
	for i, num := range nums {
		if i == index || num == 0 {
			continue
		}
		if max/num < 2 {
			isExist = false
		}
	}

	if isExist {
		return index
	}

	return index
}

// 找出最大的数和第二大的数
func DominantIndex1(nums []int) int {
	max, second := 0, 0
	idx := -1
	for i, v := range nums {
		if v > max {
			second = max
			max = v
			idx = i
		} else if v > second {
			second = v
		}
	}

	if max >= 2*second {
		return idx
	}
	return -1
}

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
输出:  [1,2,4,7,5,3,6,8,9]
*/

func FindDiagonalOrder(matrix [][]int) []int {
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
