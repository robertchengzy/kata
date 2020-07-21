package stack

type MyStack struct {
	s []int
}

/** Initialize your data structure here. */
func ConstructorS() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.s = append(this.s, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.Empty() {
		return -1
	}
	num := this.s[len(this.s)-1]
	this.s = this.s[:len(this.s)-1]
	return num
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.Empty() {
		return -1
	}
	return this.s[len(this.s)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.s) == 0
}
