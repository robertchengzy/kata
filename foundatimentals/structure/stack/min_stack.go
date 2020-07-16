package stack

/*
最小栈
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。
示例:
输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
提示：
pop、top 和 getMin 操作总是在 非空栈 上调用。
*/

type MinStack struct {
	stack []int
	mins  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: nil,
		mins:  nil,
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
	curStack := this.stack[stackLen-1]
	curMin := this.mins[minLen-1]
	this.stack = this.stack[:stackLen-1]
	if curStack == curMin {
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
	stackLen := len(this.stack)
	if stackLen == 0 {
		return 0
	}
	return this.mins[len(this.mins)-1]
}

func (this *MinStack) updateMin(x int) {
	minLen := len(this.mins)
	if minLen == 0 {
		this.mins = append(this.mins, x)
		return
	}

	cur := this.mins[minLen-1]
	if x <= cur {
		this.mins = append(this.mins, x)
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
