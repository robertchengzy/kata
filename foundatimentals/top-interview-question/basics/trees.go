package basics

import (
	"fmt"
	"strconv"
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Binary Tree Preorder Traversal 二叉树的前序遍历
/*

*/
func preorderTraversal(root *TreeNode) []int {
	pre := make([]int, 0)
	if root == nil {
		return pre
	}
	pre = append(pre, root.Val)
	pre = append(pre, preorderTraversal(root.Left)...)
	pre = append(pre, preorderTraversal(root.Right)...)
	return pre
}

func preorderTraversal2(root *TreeNode) []int {
	pre := make([]int, 0)
	preHelper(root, &pre)
	return pre
}

func preHelper(root *TreeNode, pre *[]int) {
	if root == nil {
		return
	}
	*pre = append(*pre, root.Val)
	preHelper(root.Left, pre)
	preHelper(root.Right, pre)
}

func preorderTraversal3(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		top := stack.Back()
		topNode := top.Value.(*TreeNode)
		stack.Remove(top)

		res = append(res, topNode.Val)

		if topNode.Right != nil {
			stack.PushBack(topNode.Right)
		}
		if topNode.Left != nil {
			stack.PushBack(topNode.Left)
		}
	}
	return res
}

// 先序遍历二叉树（非递归）
func preOrderNoRecursion(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		res = append(res, curNode.Val)
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
	}
	return res
}

// 中序遍历二叉树（非递归）
func inOrderNoRecursion(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	curNode := root
	for curNode != nil || len(stack) > 0 {
		// 一直循环到二叉排序树最左端的叶子结点（currentNode是null）
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}
		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(strconv.Itoa(curNode.Val) + " ")
		curNode = curNode.Right
	}
}

// 后序遍历二叉树（非递归）
func postOrderNoRecursion(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	curNode := root
	var rightNode *TreeNode
	// 一直循环到二叉排序树最左端的叶子结点（currentNode是null）
	for curNode != nil || len(stack) > 0 {
		stack = append(stack, curNode)
		curNode = curNode.Left
	}
	curNode = stack[len(stack)-1]
	// 当前结点没有右结点或上一个结点（已经输出的结点）是当前结点的右结点，则输出当前结点
	for curNode.Right == nil || curNode.Right == rightNode {
		fmt.Print(strconv.Itoa(curNode.Val) + " ")
		rightNode = curNode
		if len(stack) == 0 {
			return
		}
		curNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	stack = append(stack, curNode)
	curNode = curNode.Right
}

// 广度优先遍历二叉树，又称层次遍历二叉树
func breadthFirstRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		fmt.Print(strconv.Itoa(curNode.Val) + " ")
		if curNode.Left != nil {
			// 先将左子树入队
			queue = append(queue, curNode.Left)
		}
		if curNode.Right != nil {
			// 再将右子树入队
			queue = append(queue, curNode.Right)
		}
	}
}
