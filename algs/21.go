package algs

// 题21：https://leetcode.com/problems/merge-two-sorted-lists/
// 有序单链表 l1 和 l2，给出合并后的链表，仍要求有序

// 单链表操作题
// 时间复杂度：O(m+n) : 最好的情况是一个链表为空，另一个链表直接作为结果，复杂度 O(1)。最差的情况是 l1 和 l2 的每个节点都交叉在一起,复杂度为 O(m+n)
// 空间复杂度：O(m+n)
func Q21Solution1(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	phead := head
	for l1 != nil || l2 != nil {
		// 当其中一个链表已经没有节点了，那么可以直接把另一个链表的剩余部分直接接到 head 的后面
		if l1 == nil {
			head.Next = l2
			break
		}
		if l2 == nil {
			head.Next = l1
			break
		}

		head.Next = new(ListNode)
		if l1.Val < l2.Val {
			head.Next.Val = l1.Val
			l1 = l1.Next
		} else {
			head.Next.Val = l2.Val
			l2 = l2.Next
		}
		head = head.Next
	}

	return phead.Next
}