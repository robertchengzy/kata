package basics

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//  先序遍历二叉树（递归）
func preOrderRecursion(root *TreeNode) {
	fmt.Print(strconv.Itoa(root.Val) + " ")
	if root.Left != nil {
		preOrderRecursion(root.Left)
	}
	if root.Right != nil {
		preOrderRecursion(root.Right)
	}
}

// 中序遍历二叉树（递归）
func inOrderRecursion(root *TreeNode) {
	if root.Left != nil {
		inOrderRecursion(root.Left)
	}
	fmt.Print(strconv.Itoa(root.Val) + " ")
	if root.Right != nil {
		inOrderRecursion(root.Right)
	}
}

// 后序遍历二叉树（递归）
func postOrderRecursion(root *TreeNode) {
	if root.Left != nil {
		postOrderRecursion(root.Left)
	}
	if root.Right != nil {
		postOrderRecursion(root.Right)
	}
	fmt.Print(strconv.Itoa(root.Val) + " ")
}

// 先序遍历二叉树（非递归）
func preOrderNoRecursion(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	stack = append(stack)
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
			stack = stack[:len(stack)-1]
		}

		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
			stack = stack[:len(stack)-1]
		}
	}
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
