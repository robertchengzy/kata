package binary_tree

/*
填充每个节点的下一个右侧节点指针 II
给定一个二叉树
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有 next 指针都被设置为 NULL。
进阶：
	你只能使用常量级额外空间。
	使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
示例：
	输入：root = [1,2,3,4,5,null,7]
	输出：[1,#,2,3,#,4,5,7,#]
	解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。
提示：
	树中的节点数小于 6000
	-100 <= node.val <= 100
*/

func connectX(root *Node) *Node {
	return connect(root)
}

func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}

	leftmost = root
	curr := leftmost
	for leftmost != nil {
		prev = nil
		curr = leftmost

		leftmost = nil

		for curr != nil {
			processChild(curr.Left)
			processChild(curr.Right)

			curr = curr.Next
		}
	}
	return root
}

var prev, leftmost *Node

func processChild(childNode *Node) {
	if childNode != nil {
		if prev != nil {
			prev.Next = childNode
		} else {
			leftmost = childNode
		}

		prev = childNode
	}
}
