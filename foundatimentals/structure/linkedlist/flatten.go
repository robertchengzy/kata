package linkedlist

/*
扁平化多级双向链表
多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。
给你位于列表第一级的头节点，请你扁平化列表，使所有结点出现在单级双链表中。
示例 1：
	输入：head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
	输出：[1,2,3,7,8,11,12,9,10,4,5,6]
示例 2：
	输入：head = [1,2,null,3]
	输出：[1,3,2]
	解释：
	输入的多级列表如下图所示：
	  1---2---NULL
	  |
	  3---NULL
示例 3：
	输入：head = []
	输出：[]
如何表示测试用例中的多级链表？
以 示例 1 为例：
 1---2---3---4---5---6--NULL
         |
         7---8---9---10--NULL
             |
             11--12--NULL
序列化其中的每一级之后：
	[1,2,3,4,5,6,null]
	[7,8,9,10,null]
	[11,12,null]
为了将每一级都序列化到一起，我们需要每一级中添加值为 null 的元素，以表示没有节点连接到上一级的上级节点。
	[1,2,3,4,5,6,null]
	[null,null,7,8,9,10,null]
	[null,11,12,null]
合并所有序列化结果，并去除末尾的 null 。
	[1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]

提示：
	节点数目不超过 1000
	1 <= Node.val <= 10^5
*/

// Definition for a Node.
type FNode struct {
	Val   int
	Prev  *FNode
	Next  *FNode
	Child *FNode
}

func flatten(root *FNode) *FNode {
	if root == nil {
		return nil
	}
	pseudoHead := &FNode{
		Next: root,
	}

	flattenDFS(pseudoHead, root)

	pseudoHead.Next.Prev = nil
	return pseudoHead.Next
}

func flattenDFS(prev, curr *FNode) *FNode {
	if curr == nil {
		return prev
	}

	curr.Prev = prev
	prev.Next = curr

	tempNext := curr.Next
	tail := flattenDFS(curr, curr.Child)
	curr.Child = nil

	return flattenDFS(tail, tempNext)
}

func flatten1(root *FNode) *FNode {
	if root == nil {
		return nil
	}
	pseudoHead := &FNode{
		Next: root,
	}
	prev := pseudoHead
	stack := []*FNode{root}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		prev.Next = curr
		curr.Prev = prev
		if curr.Next != nil {
			stack = append(stack, curr.Next)
		}
		if curr.Child != nil {
			stack = append(stack, curr.Child)
			curr.Child = nil
		}
		prev = curr
	}
	pseudoHead.Next.Prev = nil
	return pseudoHead.Next
}

/*
深度优先搜索算法
递归：
·首先，我们定义递归函数 flatten_dfs(prev, curr)，它接收两个指针作为函数参数并返回扁平化列表中的尾部指针。curr 指针指向我们要扁平化的子列表，prev 指针指向 curr 指向元素的前一个元素。
·在函数 flatten_dfs(prev, curr)，我们首先在 prev 和 curr 节点之间建立双向连接。
·然后在函数中调用 flatten_dfs(curr, curr.child) 对左子树（curr.child 即子列表）进行操作，它将返回扁平化子列表的尾部元素 tail，再调用 flatten_dfs(tail, curr.next) 对右子树进行操作。
·为了得到正确的结果，我们还需要注意两个重要的细节：
	在调用 flatten_dfs(curr, curr.child) 之前我们应该复制 curr.next 指针，因为 curr.next 可能在函数中改变。
	在扁平化 curr.child 指针所指向的列表以后，我们应该删除 child 指针，因为我们最终不再需要该指针。
迭代：
关键是使用 stack 数据结构，元素遵循后进先出的原则。
stack 帮我们维持一个迭代序列，它模拟函数掉哦那个堆栈的行为，这样我们就可以不使用递归来获得相同的结果。
·首先我们创建 stack，然后将头节点压栈。利用 prev 变量帮助我们记录在每个迭代过程的前继节点。
·然后我们进入循环迭代 stack 中的元素，直到栈为空。
·在每一次迭代过程中，首先在 stack 弹出一个节点（叫做 curr）。再建立 prev 和 curr 之间的双向链接，再顺序处理 curr.next 和 curr.child 指针所指向的节点，严格按照此顺序执行。
·如果 curr.next 存在（即存在右子树），那么我们将 curr.next 压栈后进行下一次迭代。
·如果 curr.child 存在（即存在左子树），那么将 curr.child 压栈，与 curr.next 不同的是，我们需要删除 curr.child 指针，因为在最终的结果不再需要使用它。
*/
