package algs

// 题101：https://leetcode.com/problems/symmetric-tree/
// 判断一棵二叉树是否左右对称（互为镜像)。

// 这题不难，和二叉树遍历题思路一样。用递归解。
// 不过因为是要求镜像，所以要注意，需要等值的节点是 : left.left == right.right 且 left.right == right.left
// 时间复杂度：O(n) 每个节点刚好被访问一次
// 空间复杂度：O(n) ，因为一次栈空间开辟可以判断 2 个节点，所以实际上是 O(n * 1/2)，省略掉系数就是 O(n)。
func Q101Solution1(root *TreeNode) bool {
	var isMirror func(t1 *TreeNode, t2 *TreeNode) bool
	isMirror = func(t1 *TreeNode, t2 *TreeNode) bool {
		if t1 == nil && t2 == nil {
			return true
		}
		if t1 == nil || t2 == nil {
			return false
		}

		return (t1.Val == t2.Val) && isMirror(t1.Right, t2.Left) && isMirror(t1.Left, t2.Right)
	}

	return isMirror(root, root)
}