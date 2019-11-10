package algs

// 题19：https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// 单链表移除倒数第 n 个节点
// 考察点：链表结构不像数组一样可以随机访问，并且 input 是单向无循环链表，只能从 head 开始向后访问。

// 首先想到的解法就是:先遍历一遍链表，得到链长，然后移除正着数的第 length - n + 1 个元素
// 单链表操作题
// 时间复杂度：O(n) 这里的 n 不是参数 n，而是链表长度
// 空间复杂度：O(1)
func Q19Solution1(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	// 链长 listLen: 因为 head == nil 被第一步排除了，所以链表至少长度为 1
	listLen := 1
	_head := head
	for _head.Next != nil {
		listLen ++
		_head = _head.Next
	}

	// 移除倒数第 n 个元素，即正数第 listLen - n + 1 个元素
	rmNodeIndex := listLen - n + 1
	// 下面的逻辑是从 head.Next 开始，所以避免麻烦，把移除头结点这种情况先排除
	if rmNodeIndex == 1 {
		head = head.Next
		return head
	}

	// ret : 始终指向头结点, 指针不移动
	// head : 指向 current 节点
	// prev : 指向 current 节点的前一个节点
	prev := head
	ret := prev
	head = head.Next

	// 我们已经把移除头节点这种情况在之前就处理了，所以直接从第 2 个节点开始判断
	for th := 2; head != nil; th++ {
		if th == rmNodeIndex {
			prev.Next = head.Next
			break
		}
		head = head.Next
		prev = prev.Next
	}

	return ret
}