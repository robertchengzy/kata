package leetcode

import "math/rand"

// Shuffle an Array
/*
	打乱一个没有重复元素的数组。
	示例:
	// 以数字集合 1, 2 和 3 初始化数组。
	int[] nums = {1,2,3};
	Solution solution = new Solution(nums);
	// 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。
	solution.shuffle();
	// 重设数组到它的初始状态[1,2,3]。
	solution.reset();
	// 随机返回数组[1,2,3]打乱后的结果。
	solution.shuffle();
*/

type Solution struct {
	Array, Origin []int
}

func Constructor(nums []int) Solution {
	origin := make([]int, len(nums))
	array := make([]int, len(nums))
	copy(origin, nums)
	copy(array, nums)
	return Solution{Array: array, Origin: origin}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	this.Array = make([]int, len(this.Origin))
	copy(this.Array, this.Origin)
	return this.Array
}

/** Returns a random shuffling of the array. */

// Approach #1 Brute Force [Accepted]
// Time complexity : O(n^2)
// The quadratic time complexity arises from the calls to list.remove (or list.pop), which run in linear time. nn linear list removals occur, which results in a fairly easy quadratic analysis.
// Space complexity : O(n)
// Because the problem also asks us to implement reset, we must use linear additional space to store the original array. Otherwise, it would be lost upon the first call to shuffle.
func (this *Solution) Shuffle() []int {
	aux := make([]int, len(this.Array))
	copy(aux, this.Array)
	for i := 0; i < len(this.Array); i++ {
		removeIdx := rand.Intn(len(aux) - 1)
		this.Array[i] = aux[removeIdx]
		aux = append(aux[:removeIdx], aux[removeIdx+1:]...)
	}
	return this.Array
}

// Approach #2 Fisher-Yates Algorithm [Accepted]
// Time complexity : O(n)
// The Fisher-Yates algorithm runs in linear time, as generating a random index and swapping two values can be done in constant time.
// Space complexity : O(n)
// Although we managed to avoid using linear space on the auxiliary array from the brute force approach, we still need it for reset, so we're stuck with linear space complexity.
func (this *Solution) Shuffle2() []int {
	for i := range this.Array {
		index := rand.Intn(len(this.Array) - i)
		this.Array[i], this.Array[index+i] = this.Array[index+i], this.Array[i]
	}
	return this.Array
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

// Min Stack 最小栈
/*
	设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

	push(x) -- 将元素 x 推入栈中。
	pop() -- 删除栈顶的元素。
	top() -- 获取栈顶元素。
	getMin() -- 检索栈中的最小元素。
	示例:

	MinStack minStack = new MinStack();
	minStack.push(-2);
	minStack.push(0);
	minStack.push(-3);
	minStack.getMin();   --> 返回 -3.
	minStack.pop();
	minStack.top();      --> 返回 0.
	minStack.getMin();   --> 返回 -2.
*/
type MinStack struct {
	stack []int
	mins  []int
}

/** initialize your data structure here. */
func ConstructorM() MinStack {
	return MinStack{stack: nil, mins: nil}
}

func (this *MinStack) updateMin(x int) {
	minLen := len(this.mins)
	if minLen == 0 {
		this.mins = append(this.mins, x)
		return
	}
	current := this.mins[minLen-1]
	if x <= current {
		this.mins = append(this.mins, x)
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.updateMin(x)
}

func (this *MinStack) Pop() {
	stackLen := len(this.stack)
	if stackLen == 0 {
		return
	}
	minLen := len(this.mins)
	currentMin := this.mins[minLen-1]
	currentStack := this.stack[stackLen-1]
	this.stack = this.stack[:stackLen-1]
	if currentStack == currentMin {
		this.mins = this.mins[:minLen-1]
	}
}

func (this *MinStack) Top() int {
	stackLen := len(this.stack)
	if stackLen == 0 {
		return 0
	}
	return this.stack[stackLen-1]
}

func (this *MinStack) GetMin() int {
	minLen := len(this.mins)
	if minLen == 0 {
		return 0
	}
	return this.mins[minLen-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
