package algs

func Q2() map[string]func(l1 *ListNode, l2 *ListNode) *ListNode {
	solutionFuncs := map[string] func(l1 *ListNode, l2 *ListNode) *ListNode {
		"1. Linked list":Q2Solution1,
	}

	return solutionFuncs
}

// 题2：https://leetcode.com/problems/add-two-numbers/
// 单链表操作题
// 时间复杂度: O(n) 这里的 n 是 l1, l2 中的最长链的规模
// 空间复杂度: O(n) 这里的 n 是 l1, l2 中的最长链的规模
func Q2Solution1(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var dummyHead = &ListNode{Val:0}
	var p, q, curr = l1, l2, dummyHead

	for p != nil || q != nil {
		x, y := 0, 0
		if p != nil {
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}

		sum := x+y+carry
		carry = sum/10
		curr.Next = &ListNode{Val:sum%10}
		curr = curr.Next

		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}

	if carry > 0 {
		curr.Next = &ListNode{Val:carry}
	}

	return dummyHead.Next
}