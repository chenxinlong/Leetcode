package algs

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var p1 *ListNode = head.Next
	var p2 *ListNode
	head.Next = nil
	for p1 != nil {
		p2 = p1.Next
		p1.Next = head
		head = p1
		p1 = p2
	}

	return head
}