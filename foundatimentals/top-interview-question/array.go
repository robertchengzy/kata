package leetcode

// Remove Duplicates from Sorted Array
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	var index = 0
	for i := 1; i < n; i++ {
		if nums[i] != nums[index] {
			index += 1
			nums[index] = nums[i]
		}
	}

	return index + 1
}

// Best Time to Buy and Sell Stock II
func maxProfit(prices []int) int {
	var total int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			total += prices[i+1] - prices[i]
		}
	}

	return total
}
