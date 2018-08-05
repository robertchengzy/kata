package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Maximum Depth of Binary Tree  二叉树的最大深度
/*
	给定一个二叉树，找出其最大深度。

	二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

	说明: 叶子节点是指没有子节点的节点。

	示例：
	给定二叉树 [3,9,20,null,null,15,7]，

		3
	   / \
	  9  20
		/  \
	   15   7
	返回它的最大深度 3 。
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

// Validate Binary Search Tree 验证二叉搜索树
/*
	给定一个二叉树，判断其是否是一个有效的二叉搜索树。

	假设一个二叉搜索树具有如下特征：

	节点的左子树只包含小于当前节点的数。
	节点的右子树只包含大于当前节点的数。
	所有左子树和右子树自身必须也是二叉搜索树。
	示例 1:

	输入:
		2
	   / \
	  1   3
	输出: true
	示例 2:

	输入:
		5
	   / \
	  1   4
		 / \
		3   6
	输出: false
	解释: 输入为: [5,1,4,null,null,3,6]。
		 根节点的值为 5 ，但是其右子节点值为 4 。
*/
func isValidBST(root *TreeNode) bool {
	return validTree(root, math.MinInt64, math.MaxInt64)
}

func validTree(root *TreeNode, low, high int) bool {
	if root == nil {
		return true
	}

	if root.Val <= low || root.Val >= high {
		return false
	}

	return validTree(root.Left, low, root.Val) && validTree(root.Right, root.Val, high)
}

func isValidBST2(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	var pre *TreeNode

	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre != nil && root.Val <= pre.Val {
			return false
		}
		pre = root
		root = root.Right
	}

	return true
}

// Symmetric Tree 对称二叉树
/*
	给定一个二叉树，检查它是否是镜像对称的。

	例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

		1
	   / \
	  2   2
	 / \ / \
	3  4 4  3
	但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

		1
	   / \
	  2   2
	   \   \
	   3    3
	说明:

	如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
*/
// Recursive
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricHelp(root.Left, root.Right)
}

func isSymmetricHelp(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricHelp(left.Left, right.Right) && isSymmetricHelp(left.Right, right.Left)
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := make([]*TreeNode, 0)
	var left, right *TreeNode
	if root.Left != nil {
		if root.Right == nil {
			return false
		}
		stack = append(stack, root.Left, root.Right)
	} else if root.Right != nil {
		return false
	}

	for len(stack) > 0 {
		l := len(stack)
		if l%2 != 0 {
			return false
		}
		right = stack[l-1]
		left = stack[l-2]
		stack = stack[:l-2]
		if right.Val != left.Val {
			return false
		}

		if left.Left != nil {
			if right.Right == nil {
				return false
			}
			stack = append(stack, left.Left, right.Right)
		} else if right.Right != nil {
			return false
		}

		if left.Right != nil {
			if right.Left == nil {
				return false
			}
			stack = append(stack, left.Right, right.Left)
		} else if right.Left != nil {
			return false
		}
	}

	return true
}

// Binary Tree Level Order Traversal 二叉树的层次遍历
/*
	给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

	例如:
	给定二叉树: [3,9,20,null,null,15,7],

		3
	   / \
	  9  20
		/  \
	   15   7
	返回其层次遍历结果：

	[
	  [3],
	  [9,20],
	  [15,7]
	]
*/
func levelOrder(root *TreeNode) [][]int {
	return bfs(root)
}

func bfs(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	var q []*TreeNode
	q = append(q, root)
	for len(q) > 0 {
		lenQ := len(q)
		var tmp []int
		for i := 0; i < lenQ; i++ {
			n := q[i]
			tmp = append(tmp, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}

			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		q = q[lenQ:]
		ret = append(ret, tmp)
	}

	return ret
}

func levelOrder2(root *TreeNode) [][]int {
	res := make([][]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level >= len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	return res
}

// Convert Sorted Array to Binary Search Tree 将有序数组转换为二叉搜索树
/*
	将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

	本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

	示例:

	给定有序数组: [-10,-3,0,5,9],

	一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

		  0
		 / \
	   -3   9
	   /   /
	 -10  5
*/
func sortedArrayToBST(nums []int) *TreeNode {
	return bst(nums, 0, len(nums)-1)
}

func bst(nums []int, l, r int) *TreeNode {
	if l > r || l < 0 || r < 0 || l > len(nums) || r > len(nums) {
		return nil
	}

	if l == r {
		return &TreeNode{nums[l], nil, nil}
	}
	mid := (l + r) / 2
	left := bst(nums, l, mid-1)
	right := bst(nums, mid+1, r)
	return &TreeNode{nums[mid], left, right}
}

func sortedArrayToBST2(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	}

	median := len(nums) / 2
	currentNode := &TreeNode{
		Val:   nums[median],
		Left:  sortedArrayToBST2(nums[:median]),
		Right: sortedArrayToBST2(nums[median+1:]),
	}

	return currentNode
}
