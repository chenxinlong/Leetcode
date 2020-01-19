package algs

// 题102：https://leetcode.com/problems/binary-tree-level-order-traversal/
// 二叉树层次遍历

// 这题没啥好说的，按照教科书里的写就好了。
// 层次遍历，用广度优先搜索(BFS - Breadth First Search)
// 时间复杂度：O(n) 每个节点访问(queue.push/pop) 一次
// 空间复杂度：O(n)
func Q102Solution1(root *TreeNode) [][]int {
	var fnBFS func(node *TreeNode) [][]int 
	fnBFS = func(node *TreeNode) [][]int {
		if root == nil {
			return nil
		}

		var ret [][]int
		// 用 slice q 模拟一个 queue, append 模拟 queue.push,  q = q[x:] 模拟 queue.pop。
		var q []*TreeNode
		q = append(q, root)
		for len(q) > 0 {
			// 一次遍历一层，用 tmp 存放这一层所有节点

			lenQ := len(q)
			var tmp []int
			for i := 0; i < lenQ; i++ {
				n := q[i]
				tmp = append(tmp, n.Val)

				// 左子树入队列
				if n.Left != nil {
					q = append(q, n.Left)
				}
				// 右子树入队列
				if n.Right != nil {
					q = append(q, n.Right)
				}
			}
			q = q[lenQ:]
			ret = append(ret, tmp)
		}

		return ret
	}

	return fnBFS(root)
}