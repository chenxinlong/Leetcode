package algs

// 题 876：https://leetcode.com/problems/middle-of-the-linked-list/

// 第一个想到的无非就是遍历链表，节点存数组。从数组中取 middleware node 就简单多了。
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func Q876Solution1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	nodes := []*ListNode{}
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	nodesLen := len(nodes)
	return nodes[nodesLen/2]
}

// Fast slow pointer 解法(看题解的)。
// 这个思路真的碉堡了，不过也有可能是我比较菜所以想不到。 总之这个思路可以收藏一下。
// 快慢指针的解法直接把耗时降低了一半，空间复杂度直接降到 O(1)
// 从 benchmark 上看，快慢指针大概是 3~4 ns/op，而上面的第一个解法是 240 ns/op。
// 时间复杂度：O(n) 其实是 n/2，省略掉系数还是 O(n)
// 空间复杂度：O(1)
func Q876Solution2(head *ListNode) *ListNode   {
	fastPtr, slowPtr := head, head

	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}

	return slowPtr
}