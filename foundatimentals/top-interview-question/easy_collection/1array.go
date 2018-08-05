package leetcode

import (
	"sort"
)

// Remove Duplicates from Sorted Array 从排序数组中删除重复项
/*
	给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

	不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

	示例 1:

	给定数组 nums = [1,1,2],

	函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

	你不需要考虑数组中超出新长度后面的元素。
	示例 2:

	给定 nums = [0,0,1,1,1,2,2,3,3,4],

	函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

	你不需要考虑数组中超出新长度后面的元素。
	说明:

	为什么返回数值是整数，但输出的答案是数组呢?

	请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

	你可以想象内部操作如下:

	// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
	int len = removeDuplicates(nums);

	// 在函数里修改输入数组对于调用者是可见的。
	// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
	for (int i = 0; i < len; i++) {
		print(nums[i]);
	}
*/
// Time complexity : O(n). Assume that nn is the length of array. Each of i and j traverses at most n steps.
// Space complexity : O(1).
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	var index = 0
	for i := 1; i < n; i++ {
		if nums[i] != nums[index] {
			index++
			nums[index] = nums[i]
		}
	}

	return index + 1
}

func removeDuplicates2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	index := 1
	for k, v := range nums {
		if k > 0 && nums[index-1] != v {
			nums[index] = v
			index++
		}
	}
	return index
}

// Best Time to Buy and Sell Stock II 买卖股票的最佳时机 II
/*
	给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

	设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

	注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

	示例 1:

	输入: [7,1,5,3,6,4]
	输出: 7
	解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
		 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
	示例 2:

	输入: [1,2,3,4,5]
	输出: 4
	解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
		 注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
		 因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
	示例 3:

	输入: [7,6,4,3,1]
	输出: 0
	解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
*/
// Time complexity : O(n). Single pass.
// Space complexity: O(1). Constant space needed.
func maxProfit(prices []int) int {
	var total int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			total += prices[i+1] - prices[i]
		}
	}

	return total
}

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	var buy, sell, profit int
	for i := 1; i < len(prices)-1; i++ {
		if prices[i] > prices[i-1] {
			buy = prices[i-1]
		}

		if buy > 0 && prices[i] > prices[i+1] {
			sell = prices[i]
			profit += sell - buy
		}
	}

	return profit
}

// Rotate Array 旋转数组
/*
	给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

	示例 1:

	输入: [1,2,3,4,5,6,7] 和 k = 3
	输出: [5,6,7,1,2,3,4]
	解释:
	向右旋转 1 步: [7,1,2,3,4,5,6]
	向右旋转 2 步: [6,7,1,2,3,4,5]
	向右旋转 3 步: [5,6,7,1,2,3,4]
	示例 2:

	输入: [-1,-100,3,99] 和 k = 2
	输出: [3,99,-1,-100]
	解释:
	向右旋转 1 步: [99,-1,-100,3]
	向右旋转 2 步: [3,99,-1,-100]
	说明:

	尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
	要求使用空间复杂度为 O(1) 的原地算法。
*/
// Time complexity : O(n*k). All the numbers are shifted by one step(O(n)) k times(O(k)).
// Space complexity : O(1). No extra space is used.
func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		last := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = last
	}
}

// Time complexity : O(n). One pass is used to put the numbers in the new array. And another pass to copy the new array to the original one.
// Space complexity : O(n). Another array of the same size is used.
func rotate1(nums []int, k int) {
	length := len(nums)
	newNums := make([]int, length)
	for i := 0; i < length; i++ {
		newNums[(i+k)%length] = nums[i]
	}

	for i := 0; i < length; i++ {
		nums[i] = newNums[i]
	}
}

// Time complexity : O(n). Only one pass is used.
// Space complexity : O(1). Constant extra space is used.
func rotate2(nums []int, k int) {
	k = k % len(nums)
	var count = 0
	for start := 0; count < len(nums); start++ {
		current := start
		prev := nums[start]
		for {
			next := (current + k) % len(nums)
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count++

			if start == current {
				break
			}
		}
	}
}

// Time complexity : O(n). n elements are reversed a total of three times.
// Space complexity : O(1). No extra space is used.
func rotate3(nums []int, k int) {
	length := len(nums)
	k %= length
	reverse(nums, 0, length-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, length-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotate4(nums []int, k int) {
	if len(nums) == 0 {
		return
	}

	var i int
	temp := nums[0]
	start := 0
	for count := 0; count < len(nums); count++ {
		next := (i + k) % len(nums)
		nums[next], temp = temp, nums[next]

		if start == next && count+1 < len(nums) {
			start++
			i = start
			temp = nums[i]
		} else {
			i = next
		}
	}
}

// Contains Duplicate 存在重复
/*
	给定一个整数数组，判断是否存在重复元素。

	如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。

	示例 1:

	输入: [1,2,3,1]
	输出: true
	示例 2:

	输入: [1,2,3,4]
	输出: false
	示例 3:

	输入: [1,1,1,3,3,4,3,2,4,2]
	输出: true
*/
// Time complexity : O(n^2). In the worst case, there are n*(n+1)/2 pairs of integers to check. Therefore, the time complexity is O(n^2)
// Space complexity : O(1). We only used constant extra space.
func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

// Time complexity : O(n). We do search() and insert() for nn times and each operation takes constant time.
// Space complexity : O(n). The space used by a hash table is linear with the number of elements in it.
func containsDuplicate2(nums []int) bool {
	data := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := data[nums[i]]; ok {
			return true
		}
		data[nums[i]] = i
	}

	return false
}

// Time complexity : O(nlogn). Sorting is O(nlogn) and the sweeping is O(n). The entire algorithm is dominated by the sorting step, which is O(nlogn).
// Space complexity : O(1). Space depends on the sorting implementation which, usually, costs O(1) auxiliary space if heapsort is used.
func containsDuplicate3(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

// Single Number 只出现一次的数字
/*
	给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

	说明：

	你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

	示例 1:

	输入: [2,2,1]
	输出: 1
	示例 2:

	输入: [4,1,2,1,2]
	输出: 4
*/
// Time complexity : O(n⋅1)=O(n).
// Space complexity : O(n).
func singleNumber(nums []int) int {
	data := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := data[nums[i]]; ok {
			data[nums[i]] = 1
		} else {
			data[nums[i]] = 0
		}
	}

	for key := range data {
		if data[key] == 0 {
			return key
		}
	}

	return 0
}

// Time complexity : O(n). We only iterate through nums, so the time complexity is the number of elements in nums.
// Space complexity : O(1).
func singleNumber2(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v
	}
	return result
}

// Intersection of Two Arrays II 两个数组的交集 II
/*
	给定两个数组，写一个方法来计算它们的交集。

	例如:
	给定 nums1 = [1, 2, 2, 1], nums2 = [2, 2], 返回 [2, 2].

	注意：

	   输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
	   我们可以不考虑输出结果的顺序。
	跟进:

	如果给定的数组已经排好序呢？你将如何优化你的算法？
	如果 nums1 的大小比 nums2 小很多，哪种方法更优？
	如果nums2的元素存储在磁盘上，内存是有限的，你不能一次加载所有的元素到内存中，你该怎么办？
*/
func intersect(nums1 []int, nums2 []int) []int {
	dict, res := make(map[int]int), make([]int, 0, len(nums2))
	for _, v := range nums1 {
		dict[v] += 1
	}
	for _, v := range nums2 {
		if dict[v] > 0 {
			res = append(res, v)
			dict[v] -= 1
		}
	}

	return res
}

// Plus One 加一
/*
	给定一个非负整数组成的非空数组，在该数的基础上加一，返回一个新的数组。

	最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。

	你可以假设除了整数 0 之外，这个整数不会以零开头。

	示例 1:

	输入: [1,2,3]
	输出: [1,2,4]
	解释: 输入数组表示数字 123。
	示例 2:

	输入: [4,3,2,1]
	输出: [4,3,2,2]
	解释: 输入数组表示数字 4321。
*/
func plusOne(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}

	a := make([]int, length+1)
	a[0] = 1
	return a
}

// Move Zeroes 移动零
/*
	给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

	示例:

	输入: [0,1,0,3,12]
	输出: [1,3,12,0,0]
	说明:

	必须在原数组上操作，不能拷贝额外的数组。
	尽量减少操作次数。
*/
func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == 0 {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

// Time Complexity: O(n). However, the total number of operations are still sub-optimal. The total operations (array writes) that code does is nn (Total number of elements).
// Space Complexity : O(1). Only constant space is used.
func moveZeroes2(nums []int) {
	var lastNonZeroFoundAt = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZeroFoundAt] = nums[i]
			lastNonZeroFoundAt++
		}
	}

	for i := lastNonZeroFoundAt; i < len(nums); i++ {
		nums[i] = 0
	}
}

// Time Complexity: O(n). However, the total number of operations are optimal. The total operations (array writes) that code does is Number of non-0 elements.
// This gives us a much better best-case (when most of the elements are 0) complexity than last solution. However, the worst-case (when all elements are non-0) complexity for both the algorithms is same.
// Space Complexity : O(1). Only constant space is used.
func moveZeroes3(nums []int) {
	lastNonZeroFoundAt := 0
	for cur := 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[lastNonZeroFoundAt], nums[cur] = nums[cur], nums[lastNonZeroFoundAt]
			lastNonZeroFoundAt++
		}
	}
}

// Two Sum 两数之和
/*
	给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。

	你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

	示例:

	给定 nums = [2, 7, 11, 15], target = 9

	因为 nums[0] + nums[1] = 2 + 7 = 9
	所以返回 [0, 1]
*/
// Time complexity : O(n^2). For each element, we try to find its complement by looping through the rest of array which takes O(n) time. Therefore, the time complexity is O(n^2).
//Space complexity : O(1).
func twoSum(nums []int, target int) []int {
	res := make([]int, 0, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)
				return res
			}
		}
	}

	return res
}

func twoSum2(nums []int, target int) []int {
	data := map[int]int{}
	for idx, v := range nums {
		res := target - v
		if idxSec, ok := data[res]; ok {
			return []int{idxSec, idx}
		}
		data[v] = idx
	}

	return nil
}

// Valid Sudoku 有效的数独
/*
	判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

	数字 1-9 在每一行只能出现一次。
	数字 1-9 在每一列只能出现一次。
	数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。


	上图是一个部分填充的有效的数独。

	数独部分空格内已填入了数字，空白格用 '.' 表示。

	示例 1:

	输入:
	[
	  ["5","3",".",".","7",".",".",".","."],
	  ["6",".",".","1","9","5",".",".","."],
	  [".","9","8",".",".",".",".","6","."],
	  ["8",".",".",".","6",".",".",".","3"],
	  ["4",".",".","8",".","3",".",".","1"],
	  ["7",".",".",".","2",".",".",".","6"],
	  [".","6",".",".",".",".","2","8","."],
	  [".",".",".","4","1","9",".",".","5"],
	  [".",".",".",".","8",".",".","7","9"]
	]
	输出: true
	示例 2:

	输入:
	[
	  ["8","3",".",".","7",".",".",".","."],
	  ["6",".",".","1","9","5",".",".","."],
	  [".","9","8",".",".",".",".","6","."],
	  ["8",".",".",".","6",".",".",".","3"],
	  ["4",".",".","8",".","3",".",".","1"],
	  ["7",".",".",".","2",".",".",".","6"],
	  [".","6",".",".",".",".","2","8","."],
	  [".",".",".","4","1","9",".",".","5"],
	  [".",".",".",".","8",".",".","7","9"]
	]
	输出: false
	解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
		 但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。
	说明:

	一个有效的数独（部分已被填充）不一定是可解的。
	只需要根据以上规则，验证已经填入的数字是否有效即可。
	给定数独序列只包含数字 1-9 和字符 '.' 。
	给定数独永远是 9x9 形式的。
*/
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		dataRow := make(map[byte]int)
		dataLine := make(map[byte]int)
		for j := 0; j < 9; j++ {
			if string(board[i][j]) != "." {
				if _, ok := dataRow[board[i][j]]; ok {
					return false
				}
				dataRow[board[i][j]] = 1
			}

			if string(board[j][i]) != "." {
				if _, ok := dataLine[board[j][i]]; ok {
					return false
				}
				dataLine[board[j][i]] = 1
			}
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			data := make(map[byte]int)
			for m := i; m < i+3; m++ {
				for n := j; n < j+3; n++ {
					if string(board[m][n]) != "." {
						if _, ok := data[board[m][n]]; ok {
							return false
						}
						data[board[m][n]] = 1
					}
				}
			}
		}
	}

	return true
}

func isValidSudoku2(board [][]byte) bool {
	colUsed := [9][9]bool{}
	subBoxUsed := [3][3][9]bool{}

	for row := 0; row < 9; row++ {
		rowUsed := [9]bool{}
		for col := 0; col < 9; col++ {
			b := board[row][col]
			if b == '.' {
				continue
			}

			num := b - '0' - 1

			if rowUsed[num] || colUsed[col][num] || subBoxUsed[row/3][col/3][num] {
				return false
			}

			rowUsed[num] = true
			colUsed[col][num] = true
			subBoxUsed[row/3][col/3][num] = true
		}
	}

	return true
}

// Rotate Image 旋转图像
/*
	给定一个 n × n 的二维矩阵表示一个图像。

	将图像顺时针旋转 90 度。

	说明：

	你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

	示例 1:

	给定 matrix =
	[
	  [1,2,3],
	  [4,5,6],
	  [7,8,9]
	],

	原地旋转输入矩阵，使其变为:
	[
	  [7,4,1],
	  [8,5,2],
	  [9,6,3]
	]
	示例 2:

	给定 matrix =
	[
	  [ 5, 1, 9,11],
	  [ 2, 4, 8,10],
	  [13, 3, 6, 7],
	  [15,14,12,16]
	],

	原地旋转输入矩阵，使其变为:
	[
	  [15,13, 2, 5],
	  [14, 3, 4, 1],
	  [12, 6, 8, 9],
	  [16, 7,10,11]
	]
*/
func rotateMatrix(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l/2; j++ {
			matrix[i][j], matrix[i][l-1-j] = matrix[i][l-1-j], matrix[i][j]
		}
	}
}

func rotateMatrix2(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < l/2; i++ {
		for j := i; j < l-i-1; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[l-j-1][i]
			matrix[l-j-1][i] = matrix[l-i-1][l-j-1]
			matrix[l-i-1][l-j-1] = matrix[j][l-i-1]
			matrix[j][l-i-1] = temp
		}
	}
}
