package algs

import "sort"

// 题23：https://leetcode.com/problems/merge-k-sorted-lists/
// 合并多个有序链表，要求合并后的链表仍然有序。

// 像我这种不会算法的，第一个想到的就是 : 所有节点值存在数组里，排序一下再构造出结果。
// 时间复杂度：O(k*n + m) k 是链表数量，n 是链表平均节点数，m 是排序所需时间复杂度。
// 空间复杂度：O(k*n) k 是链表数量，n 是链表平均节点数。
func Q23Soution1(lists []*ListNode) *ListNode {
	head := &ListNode{}
	p := head

	allNodeValues := []int{}
	listsNum := len(lists)
	for i:=0; i < listsNum; i++ {
		for lists[i] != nil {
			allNodeValues = append(allNodeValues, lists[i].Val)
			lists[i] = lists[i].Next
		}
	}
	sort.Ints(allNodeValues)

	for _, val := range allNodeValues {
		node := &ListNode{
			Val:val,
		}
		p.Next = node
		p = p.Next
	}

	return head.Next
}

// 用 reduce 的思路, 每次仅合并2个链表并产生结果链表(合并2个链表的方法见 21 题)，用结果链表与下一个链表再合并。
// 这种解法看起来比第 1 种解法高明一丢丢，但是实际上效率比第一种垃圾一大截。当然, reduce 的产出有点慢，需要合并 k-1 次，也可以用双指针继续优化。
// 时间复杂度：O(kn) k 是 lists 长度，n 是链表平均节点数。实际上是 O(k(m+n))，这里的 m+n 可以用 2n表示，即 O(2kn)，去掉系数之后是 O(kn)。
// 空间复杂度：O(kn) k 是 lists 长度，n 是链表平均节点数。因为我们重新构造了一个新链表，所以空间复杂度就是所有节点个数之和。
func Q23Solution2(lists []*ListNode) *ListNode {
	listsNum := len(lists)
	if listsNum < 2 {
		return lists[0]
	}

	// 时间复杂度：O(m+n)
	// 空间复杂度：O(m+n)
	fnMergeTwoList := func(a, b *ListNode) *ListNode {
		head := &ListNode{}
		p := head
		for a != nil || b != nil {
			if a == nil {
				p.Next = b
				break
			}
			if b == nil {
				p.Next = a
				break
			}

			p.Next = new(ListNode)
			if a.Val < b.Val {
				p.Next.Val = a.Val
				a = a.Next
			} else {
				p.Next.Val = b.Val
				b = b.Next
			}
			p = p.Next
		}
		return head.Next
	}

	var mergedLinkList *ListNode
	for i:=0; i<listsNum; i++ {
		mergedLinkList = fnMergeTwoList(mergedLinkList, lists[i])
	}

	return mergedLinkList
}