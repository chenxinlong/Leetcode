package algs
// 题92：https://leetcode.com/problems/reverse-linked-list-ii/
// 翻转链表题，加了个条件，是翻转 m~n 节点之间的那段链。

// 时间复杂度：O(m-n)
// 空间复杂度：O(n) 这里的 n 指的是链表长度，不是参数 n。
func Q92Solution1(head *ListNode, m int, n int) *ListNode {
	result := &ListNode{0, head}

	left := result
	for i := 0; i < m-1; i++ {
		left = left.Next
	}

	middle := left.Next
	right := middle.Next

	for i := 0; i < n-m; i++ {
		middle.Next = right.Next
		right.Next = left.Next
		left.Next = right
		right = middle.Next
	}

	return result.Next
}