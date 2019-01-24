package algs


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsUnivalTree(root *TreeNode) bool {
	return rec(root.Left, root.Val) && rec(root.Right, root.Val)
}


func rec(node *TreeNode, firstVal int) bool {
	if node == nil {
		return true
	}

	if node.Val == firstVal {
		return rec(node.Left, firstVal) && rec(node.Right, firstVal)
	}

	return false
}

/**
 * [analysis]
 * use recursion
 */
