package algs

// 题94：https://leetcode.com/problems/binary-tree-inorder-traversal/
// 二叉树中序遍历

// 二叉树基础题，前、中、后序遍历。 同时也是递归的经典题。
// 时间复杂度：O(n) 因为每个节点刚好被 append 一次
// 空间复杂度：O(n) 因为每个节点一个栈空间
func Q94Solution1(root *TreeNode) []int {
	var traversal func(root *TreeNode, result []int) []int
	traversal = func(root *TreeNode, result []int) []int {
		if root != nil {
			// 中序遍历 :  "左 根 右"
			if root.Left != nil {
				result = traversal(root.Left, result)
			}
			result = append(result, root.Val)
			if root.Right != nil {
				result = traversal(root.Right, result)
			}
		}

		return result
	}

	result := []int{}
	if root == nil {
		return result
	}

	return  traversal(root, []int{})
}