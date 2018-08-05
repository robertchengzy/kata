package leetcode

import (
	"sort"
)

// Remove Duplicates from Sorted Array 从排序数组中删除重复项
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
