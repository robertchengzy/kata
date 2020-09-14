package linkedlist

/*
复制带随机指针的链表
给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。
要求返回这个链表的 深拷贝。
我们用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：
	·val：一个表示 Node.val 的整数。
	·random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为 null 。
示例 1：
	输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
	输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
示例 2：
	输入：head = [[1,1],[2,1]]
	输出：[[1,1],[2,1]]
示例 3：
	输入：head = [[3,null],[3,0],[3,null]]
	输出：[[3,null],[3,0],[3,null]]
示例 4：
	输入：head = []
	输出：[]
	解释：给定的链表为空（空指针），因此返回 null。
提示：
	-10000 <= Node.val <= 10000
	Node.random 为空（null）或指向链表中的节点。
	节点数目不超过 1000 。
*/

// Definition for a Node.
type RNode struct {
	Val    int
	Next   *RNode
	Random *RNode
}

// 回溯
var visitedMap = make(map[*RNode]*RNode)

func copyRandomList(head *RNode) *RNode {
	if head == nil {
		return nil
	}

	node, ok := visitedMap[head]
	if ok {
		return node
	}

	node = &RNode{
		Val: head.Val,
	}

	visitedMap[head] = node

	node.Next = copyRandomList(head.Next)
	node.Random = copyRandomList(head.Random)

	return node
}

// 空间迭代
func copyRandomList1(head *RNode) *RNode {
	if head == nil {
		return nil
	}

	oldNode := head
	newNode := &RNode{
		Val: oldNode.Val,
	}
	visitedMap[oldNode] = newNode
	for oldNode != nil {
		newNode.Random = getClonedNode(oldNode.Random)
		newNode.Next = getClonedNode(oldNode.Next)

		oldNode = oldNode.Next
		newNode = newNode.Next
	}

	return visitedMap[head]
}

func getClonedNode(node *RNode) *RNode {
	if node != nil {
		curr, ok := visitedMap[node]
		if ok {
			return curr
		} else {
			visitedMap[node] = &RNode{
				Val: node.Val,
			}
		}
	}
	return nil
}

// 空间的迭代
func copyRandomList2(head *RNode) *RNode {
	if head == nil {
		return nil
	}
	ptr := head
	for ptr != nil {
		newNode := &RNode{
			Val: ptr.Val,
		}
		newNode.Next = ptr.Next
		ptr.Next = newNode
		ptr = newNode.Next
	}

	ptr = head
	for ptr != nil {
		if ptr.Random != nil {
			ptr.Next.Random = ptr.Random.Next
		} else {
			ptr.Next.Random = nil
		}
		ptr = ptr.Next.Next
	}

	ptrOldList := head
	ptrNewList := head.Next
	headOld := head.Next
	for ptrOldList != nil {
		ptrOldList.Next = ptrOldList.Next.Next
		if ptrNewList.Next != nil {
			ptrNewList.Next = ptrNewList.Next.Next
		} else {
			ptrNewList.Next = nil
		}
		ptrOldList = ptrOldList.Next
		ptrNewList = ptrNewList.Next
	}
	return headOld
}
