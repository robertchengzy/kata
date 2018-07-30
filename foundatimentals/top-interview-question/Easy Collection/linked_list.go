package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Delete Node in a Linked List
// Time and space complexity are both O(1).
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// Remove Nth Node From End of List
// Time complexity : O(L).
// The algorithm makes two traversal of the list, first to calculate list length LL and second to find the (L−n) th node. There are 2L-n operations and time complexity is O(L).
// Space complexity : O(1). We only used constant extra space.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Val = 0
	dummy.Next = head
	length := 0

	first := head
	for first != nil {
		length++
		first = first.Next
	}

	length -= n
	first = dummy
	for length > 0 {
		length--
		first = first.Next
	}

	first.Next = first.Next.Next
	return dummy.Next
}

// Time complexity : O(L). The algorithm makes one traversal of the list of LL nodes. Therefore time complexity is O(L).
// Space complexity : O(1). We only used constant extra space.
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Val = 0
	dummy.Next = head

	first, second := dummy, dummy
	// Advances first pointer so that the gap between first and second is n nodes apart
	for i := 1; i <= n+1; i++ {
		first = first.Next
	}

	// Move first to the end, maintaining the gap
	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummy.Next
}

// Reverse Linked List
// Time complexity : O(n). Assume that nn is the list's length, the time complexity is O(n).
// Space complexity : O(1).
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		nextTmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTmp
	}

	return prev
}

// Time complexity : O(n). Assume that nn is the list's length, the time complexity is O(n).
// Space complexity : O(n). The extra space comes from implicit stack space due to recursion. The recursion could go up to nn levels deep.
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

// Merge Two Sorted Lists
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

// Palindrome Linked List
func isPalindromeLinkedList(head *ListNode) bool {
	var slowNext, slowPrev *ListNode
	slow := head
	fast := head

	// Reverse the first half of the linked list
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		slowNext = slow.Next
		slow.Next = slowPrev
		slowPrev = slow
		slow = slowNext
	}

	if fast != nil {
		slow = slow.Next
	}

	// Next compare the two halves and check for equality
	for slowPrev != nil && slowPrev.Val == slow.Val {
		slowPrev = slowPrev.Next
		slow = slow.Next
	}

	return slowPrev == nil
}

// Linked List Cycle
// 1.Use two pointers, walker and runner.
// 2.walker moves step by step. runner moves two steps at time.
// 3.if the Linked List has a cycle walker and runner will meet at some point.
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	walker, runner := head, head
	for runner.Next != nil && runner.Next.Next != nil {
		walker = walker.Next
		runner = runner.Next.Next
		if walker == runner {
			return true
		}
	}

	return false
}
