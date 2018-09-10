package algs

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)

	if left >= right {
		return left +1
	}

	return right +1
}

/**
 * [analysis]
 * use recursion
 */