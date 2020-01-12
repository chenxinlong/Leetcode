package algs

// 题141：https://leetcode.com/problems/linked-list-cycle/
// 这题出的有点问题，从 leetcode 给出的方法签名来看，参数里并没有 pos。所以看了下 discussion，大家普遍都是直接判断链表是否成环。

// 判断是否成环其实就很简单了，只要某个节点的 next 节点指向该节点之前的节点，就是成环。
// 这样一来用 map 处理就好了。
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func Q141Solution1(head *ListNode) bool {
	m := map[*ListNode]bool{}
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = true
		head = head.Next
	}

	return false
}