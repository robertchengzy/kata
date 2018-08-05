package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Maximum Depth of Binary Tree
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

// Validate Binary Search Tree
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

// Symmetric Tree
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

// Binary Tree Level Order Traversal
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

// Convert Sorted Array to Binary Search Tree
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
