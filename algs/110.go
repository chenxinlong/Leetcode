package algs

// 题110：https://leetcode.com/problems/balanced-binary-tree/
// 判断一棵二叉树是否平衡

// 二叉树，递归就完事儿了。
// 首先你了解什么是平衡二叉树(AVL树) : 任一节点的左右子树深度之差不超过1即平衡。
// 了解概念我们就知道把所以节点的左右子树 depth 求出来再比较即可。
// 时间复杂度：O(n) n 为树节点数。
// 空间复杂度：O(1)
func Q110Solution1(root *TreeNode) bool {
	isBalance := true

	var fnMaxDepth func(node *TreeNode) int
	fnMaxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftChildDepth := fnMaxDepth(node.Left)
		rightChildDepth := fnMaxDepth(node.Right)
		if leftChildDepth - rightChildDepth > 1 || rightChildDepth -leftChildDepth > 1 {
			isBalance = false
		}

		if leftChildDepth > rightChildDepth {
			return 1 + leftChildDepth
		}

		return 1 + rightChildDepth
	}

	fnMaxDepth(root)
	return isBalance
}