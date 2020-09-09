package binary_tree

/*
从前序与中序遍历序列构造二叉树
根据一棵树的前序遍历与中序遍历构造二叉树。
注意:
你可以假设树中没有重复的元素。
例如，给出
	前序遍历 preorder = [3,9,20,15,7]
	中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
*/

var (
	idxMapX   = make(map[int]int)
	preorderx []int
	preIdx    int
)

func buildTreeX(preorder []int, inorder []int) *TreeNode {
	preorderx = preorder
	preIdx = 0

	idx := 0
	for _, val := range inorder {
		idxMapX[val] = idx
		idx++
	}
	return helperX(0, len(inorder)-1)
}

func helperX(inLeft, inRight int) *TreeNode {
	if inLeft > inRight {
		return nil
	}

	rootVal := preorderx[preIdx]
	root := &TreeNode{
		Val: rootVal,
	}

	index := idxMapX[rootVal]
	preIdx++
	root.Left = helperX(inLeft, index-1)
	root.Right = helperX(index+1, inRight)
	return root
}

func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	var stack []*TreeNode
	stack = append(stack, root)
	var inorderIndex int
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stack = append(stack, node.Right)
		}
	}
	return root
}
