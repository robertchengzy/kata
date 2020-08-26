package stack

/*
二叉树的中序遍历
给定一个二叉树，返回它的中序遍历。
示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3
输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
1.前序遍历：先访问根节点——左子树——右子树。
2.中序遍历：先访问左子树——根节点——右子树，按照这个顺序。
3.后序遍历：和前面差不多，先访问树的左子树——右子树——根节点。按层遍历：把一棵树从上到下，从左到右依次写出来。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res
}

func inorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	curNode := root
	for curNode != nil || len(stack) > 0 {
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}
		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curNode.Val)
		curNode = curNode.Right
	}
	return res
}