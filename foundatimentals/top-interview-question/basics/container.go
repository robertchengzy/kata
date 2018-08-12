package basics

type Node struct {
	data interface{}
	next *Node
}

type Queue struct {
	head *Node
	end  *Node
}

func NewQueue() *Queue {
	q := &Queue{nil, nil}
	return q
}

func (q *Queue) push(data interface{}) {
	n := &Node{data: data, next: nil}

	if q.end == nil {
		q.head = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}

	return
}

func (q *Queue) pop() (interface{}, bool) {
	if q.head == nil {
		return nil, false
	}

	data := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.end = nil
	}

	return data, true
}

type Stack struct {
	head *Node
}

func NewStack() *Stack {
	s := &Stack{nil}
	return s
}

func (s *Stack) Push(data interface{}) {
	n := &Node{data: data, next: s.head}
	s.head = n
}

func (s *Stack) Pop() (interface{}, bool) {
	n := s.head
	if s.head == nil {
		return nil, false
	}
	s.head = s.head.next
	return n.data, true
}
