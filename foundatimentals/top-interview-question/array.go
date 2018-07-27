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

func removeDuplicates2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	length := 1
	for k, v := range nums {
		if k > 0 && nums[length-1] != v {
			nums[length] = v
			length++
		}
	}
	return length
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

// Rotate Array
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
