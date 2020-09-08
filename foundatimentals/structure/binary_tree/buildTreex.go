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
