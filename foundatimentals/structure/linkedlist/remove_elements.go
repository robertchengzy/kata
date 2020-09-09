package linkedlist

/*
移除链表元素
删除链表中等于给定值 val 的所有节点。
示例:
	输入: 1->2->6->3->4->5->6, val = 6
	输出: 1->2->3->4->5
*/

func removeElements(head *ListNode, val int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	first := dummy
	second := dummy.Next
	for second != nil {
		if second.Val == val {
			first.Next = second.Next
		} else {
			first = first.Next
		}
		second = second.Next
	}
	return dummy.Next
}
