package queue

type Node struct {
	value int
	next  *Node
}

type MyCircularQueueX struct {
	head, tail      *Node
	count, capacity int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func ConstructorX(k int) MyCircularQueueX {
	return MyCircularQueueX{
		head:     nil,
		tail:     nil,
		count:    0,
		capacity: k,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueueX) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	node := &Node{
		value: value,
		next:  nil,
	}
	if this.count == 0 {
		this.head = node
		this.tail = node
	} else {
		this.tail.next = node
		this.tail = node
	}

	this.count++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueueX) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = this.head.next
	this.count--
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueueX) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.head.value
}

/** Get the last item from the queue. */
func (this *MyCircularQueueX) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.tail.value
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueueX) IsEmpty() bool {
	return this.count == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueueX) IsFull() bool {
	return this.count == this.capacity
}
