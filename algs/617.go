package algs

func MergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	return &TreeNode{Val: t1.Val+t2.Val, Left:MergeTrees(t1.Left, t2.Left), Right:MergeTrees(t1.Right, t2.Right)}
}

/**
 * [analysis]
 * Binary often use recursion
 * Time complexity : depends on the depth of the deepest tree
 */