package algs

// 题234：https://leetcode.com/problems/palindrome-linked-list/

// 第一个想到的还是存到数组,然后用头尾双指针判断数组是否回文。
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func Q234Solution1(head *ListNode) bool {
	if head == nil {
		return true
	}

	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	for left,right := 0, len(arr)-1; left <= right; left, right = left+1, right-1 {
		if arr[left] != arr[right] {
			return false
		}
	}
	return true
}

// 快慢指针 + 链表翻转
// 初始阶段：[1 -> 2 -> 3 -> 3 -> 2 -> 1]
// 翻转完成：[3 -> 2 -> 1 3 -> 2 -> 1]
//          ↑           ↑
//         left       right
//
// 翻转完前半个链表，此时只要逐个对比 left 和 right 的值是否相等就行了。
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func Q234Solution2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 快慢指针
	slowPtr, fastPtr := head, head

	// 用于翻转链表的指针
	var p1, p2 *ListNode

	// 通过快慢指针，让链表的前半部分翻转。 同时得到一个指向 middle 节点的指针。
	// 快指针的唯一作用就是触尾后让慢指针到达链表的 middle 节点。
	for {
		fastPtr = fastPtr.Next.Next

		// 翻转慢指针经过的节点
		p1 = slowPtr
		slowPtr = slowPtr.Next
		p1.Next = p2
		p2 = p1

		if fastPtr == nil || fastPtr.Next == nil {
			break
		}
	}

	// 此时慢指针指向 middle 节点
	left, right := p1, slowPtr
	if fastPtr != nil {
		right = right.Next
	}
	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}