package linkedlist

/*
删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
示例：
	给定一个链表: 1->2->3->4->5, 和 n = 2.
	当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：
给定的 n 保证是有效的。
进阶：
你能尝试使用一趟扫描实现吗？

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
时间复杂度：O(L)，该算法对含有 L 个结点的列表进行了一次遍历。因此时间复杂度为 O(L)。
空间复杂度：O(1)，我们只用了常量级的额外空间。
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	start := dummy
	end := dummy
	for i := 0; i < n+1; i++ {
		end = end.Next
	}

	for end != nil {
		end = end.Next
		start = start.Next
	}

	start.Next = start.Next.Next
	return dummy.Next
}
