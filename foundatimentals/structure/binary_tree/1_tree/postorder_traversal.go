package tree

/*
二叉树的后序遍历
给定一个二叉树，返回它的 后序 遍历。
示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3
输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return nil
	}
	res = append(res, preorderTraversal(root.Right)...)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, root.Val)
	return res
}

func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}

	var ans []int
	l := len(res)
	for i := l - 1; i >= 0; i-- {
		ans = append(ans, res[i])
	}
	return ans
}
