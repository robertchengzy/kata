package binary_tree

/*
从中序与后序遍历序列构造二叉树
根据一棵树的中序遍历与后序遍历构造二叉树。
注意:
你可以假设树中没有重复的元素。
例如，给出
中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	postorderx []int
	postIdx    int
	idxMap     = make(map[int]int)
)

func buildTree(inorder []int, postorder []int) *TreeNode {
	postorderx = postorder
	// start from the last postorder element
	postIdx = len(postorder) - 1

	// build a hashmap value -> its index
	idx := 0
	for _, val := range inorder {
		idxMap[val] = idx
		idx++
	}
	return helper(0, len(inorder)-1)
}

func helper(inLeft, inRight int) *TreeNode {
	// if there is no elements to construct subtrees
	if inLeft > inRight {
		return nil
	}

	// pick up post_idx element as a root
	rootVal := postorderx[postIdx]
	root := &TreeNode{
		Val: rootVal,
	}

	// root splits inorder list
	// into left and right subtrees
	index := idxMap[rootVal]

	// recursion
	postIdx--

	// build right subtree
	root.Right = helper(index+1, inRight)
	// build left subtree
	root.Left = helper(inLeft, index-1)
	return root
}

/*
算法

创建 hashmap 存储中序序列：value -> its index 。

方法 helper 的参数是中序序列中当前子树的左右边界，该方法仅用于检查子树是否为空。下面分析 helper(in_left = 0, in_right = n - 1) 的逻辑：

	如果 in_left > in_right，说明子树为空，返回 None。

	选择后序遍历的最后一个节点作为根节点。

	假设根节点在中序遍历中索引为 index。从 in_left 到 index - 1 属于左子树，从 index + 1 到 in_right 属于右子树。

	根据后序遍历逻辑，递归创建右子树 helper(index + 1, in_right) 和左子树 helper(in_left, index - 1)。

	返回根节点 root。
*/
