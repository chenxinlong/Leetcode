package algs

// 一些用于简化流程的 helper 方法，里的方法应该与算法的规模 n 无关, 只能是一些 O(1) 的简单流程封装.
//
// func max
// func min
// 因为 Go 1.x 中暂时没有三元运算符、运算符重载、泛型。
// 如果要实现一个简单的数字比较，三元运算符只需要 max = a > b ? a : b ，如果用 if else 多写好几行代码
// 用 max 只需要 max = max(a,b)

// 返回 a 与 b 中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 返回 a 与 b 中的较小值
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 返回 a 与 b 中的较大值
func maxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// 返回 a 与 b 中的较小值
func minFloat64(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}

// 数组转单链表
func ArrayToSingleLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := new(ListNode)
	p := head
	for _, n := range arr {
		p.Next = &ListNode{
			Val:n,
		}
		p = p.Next
	}
	return head.Next
}

// 单链表转数组
func SingleLinkedListToArray(head *ListNode) []int {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return arr
}