package main

import "fmt"

type Node struct {
	data Object
	next *Node
}

type Object interface{}

type List struct {
	size uint64
	head *Node
	tail *Node
}

func (list *List) Init() {
	list.size = 0
	list.head = nil
	list.tail = nil
}

func (list *List) Append(node *Node) bool {
	if node == nil {
		return false
	}

	node.next = nil
	if list.size == 0 {
		list.head = node
	} else {
		list.tail.next = node
	}

	list.tail = node
	list.size++

	return true
}

func main() {
	list := new(List)
	list.Init()
	node := &Node{
		data: "123",
		next: nil,
	}
	list.Append(node)
	node2 := &Node{
		data: "1234",
		next: nil,
	}
	list.Append(node2)
	fmt.Println(list.head)
	fmt.Println(list.head.next)
}
