package linkedlist

/*
时间复杂度：
	addAtHead，addAtTail： O(1)
	get，addAtIndex，delete：O(min(k,N−k))，其中 k 指的是元素的索引。
空间复杂度：所有的操作都是 O(1)。
*/

type Node struct {
	Val        int
	Prev, Next *Node
}

type MyLinkedListX struct {
	Size       int
	Head, Tail *Node
}

/** Initialize your data structure here. */
func ConstructorX() MyLinkedListX {
	head := new(Node)
	tail := new(Node)
	head.Next = tail
	tail.Prev = head
	return MyLinkedListX{
		Head: head,
		Tail: tail,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedListX) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}
	cur := this.Head
	if index+1 < this.Size-index {
		for i := 0; i < index+1; i++ {
			cur = cur.Next
		}
	} else {
		cur = this.Tail
		for i := 0; i < this.Size-index; i++ {
			cur = cur.Prev
		}
	}

	return cur.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedListX) AddAtHead(val int) {
	pred, succ := this.Head, this.Head.Next

	node := &Node{Val: val}
	node.Prev = pred
	node.Next = succ
	pred.Next = node
	succ.Prev = node
	this.Size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedListX) AddAtTail(val int) {
	succ, pred := this.Tail, this.Tail.Prev

	node := &Node{Val: val}
	node.Prev = pred
	node.Next = succ
	pred.Next = node
	succ.Prev = node
	this.Size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedListX) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	}
	if index < 0 {
		index = 0
	}
	var pred, succ *Node
	if index < this.Size-index {
		pred = this.Head
		for i := 0; i < index; i++ {
			pred = pred.Next
		}
		succ = pred.Next
	} else {
		succ = this.Tail
		for i := 0; i < this.Size-index; i++ {
			succ = succ.Prev
		}
		pred = succ.Prev
	}

	node := &Node{Val: val}
	node.Prev = pred
	node.Next = succ
	pred.Next = node
	succ.Prev = node
	this.Size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedListX) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	var pred, succ *Node
	if index < this.Size-index {
		pred = this.Head
		for i := 0; i < index; i++ {
			pred = pred.Next
		}
		succ = pred.Next.Next
	} else {
		succ = this.Tail
		for i := 0; i < this.Size-index-1; i++ {
			succ = succ.Prev
		}
		pred = succ.Prev.Prev
	}

	pred.Next = succ
	succ.Prev = pred
	this.Size--
}
