package recursive

import "kata/foundatimentals/structure/binary_tree/tree"

/*
二叉树的最大深度
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。
示例：
给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *tree.TreeNode) int {
	answer = 0
	maximumDepth(root, 1)
	return answer
}

var answer int

func maximumDepth(root *tree.TreeNode, depth int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		answer = max(answer, depth)
	}
	maximumDepth(root.Left, depth+1)
	maximumDepth(root.Right, depth+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth1(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth1(root.Left), maxDepth1(root.Right)) + 1
}

func maxDepth2(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*tree.TreeNode
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}
