package algs

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	tmp := InvertTree(root.Left)
	root.Left = InvertTree(root.Right)
	root.Right = tmp

	return root
}

/**
 * [analysis]
 * use recursion
 */