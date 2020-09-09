package linkedlist

/*
回文链表
请判断一个链表是否为回文链表。
示例 1:
	输入: 1->2
	输出: false
示例 2:
	输入: 1->2->2->1
	输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/

/*
算法：
	复制链表值到数组列表中。
	使用双指针法判断是否为回文。
时间复杂度：O(n)，其中 n 指的是链表的元素个数。
	第一步： 遍历链表并将值复制到数组中，O(n)。
	第二步：双指针判断是否为回文，执行了 O(n/2) 次的判断，即 O(n)。
	总的时间复杂度：O(2n) = O(n)。
空间复杂度：O(n)，其中 n 指的是链表的元素个数，我们使用了一个数组列表存放链表的元素值。
*/
func isPalindrome(head *ListNode) bool {
	var vals []int
	cur := head
	for cur != nil {
		vals = append(vals, cur.Val)
		cur = cur.Next
	}

	front := 0
	back := len(vals) - 1
	for front < back {
		if vals[front] != vals[back] {
			return false
		}
		front++
		back--
	}
	return true
}

/*
算法：
	currentNode 指针是先到尾节点，由于递归的特性再从后往前进行比较。frontPointer 是递归函数外的指针。
	若 currentNode.val != frontPointer.val 则返回 false。反之，frontPointer 向前移动并返回 true。
时间复杂度：O(n)，其中 n 指的是链表的大小。
空间复杂度：O(n)，其中 n 指的是链表的大小。我们要理解计算机如何运行递归函数，在一个函数中调用一个函数时，
	计算机需要在进入被调用函数之前跟踪它在当前函数中的位置（以及任何局部变量的值），通过运行时存放在堆栈中来实现（堆栈帧）。
	在堆栈中存放好了数据后就可以进入被调用的函数。在完成被调用函数之后，他会弹出堆栈顶部元素，以恢复在进行函数调用之前所在的函数。
	在进行回文检查之前，递归函数将在堆栈中创建 n 个堆栈帧，计算机会逐个弹出进行处理。所以在使用递归时要考虑堆栈的使用情况。
*/
var frontPointer *ListNode

func recursivelyCheck(cur *ListNode) bool {
	if !recursivelyCheck(cur.Next) {
		return false
	}
	if cur.Val != frontPointer.Val {
		return false
	}
	frontPointer = frontPointer.Next
	return true
}

func isPalindrome2(head *ListNode) bool {
	frontPointer = head
	return recursivelyCheck(head)
}

/*
算法：
	找到前半部分链表的尾节点。
	反转后半部分链表。
	判断是否为回文。
	恢复链表。
	返回结果。
时间复杂度：O(n)，其中 n 指的是链表的大小。
空间复杂度：O(1)，我们是一个接着一个的改变指针，我们在堆栈上的堆栈帧不超过 O(1)。
该方法的缺点是，在并发环境下，函数运行时需要锁定其他线程或进程对链表的访问，因为在函数执执行过程中链表暂时断开。
*/
func isPalindrome3(head *ListNode) bool {
	if head == nil {
		return true
	}

	// Find the end of first half and reverse second half.
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	// Check whether or not there is a palindrome.
	p1, p2 := head, secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// Restore the list and return the result.
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}

func endOfFirstHalf(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
