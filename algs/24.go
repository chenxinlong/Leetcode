package algs

// 题24：https://leetcode.com/problems/swap-nodes-in-pairs/

// 这题感觉没有啥难度，但是不知道为啥通过率还不到一半。
// 漏了一个边界值判断，所以失败了一次。
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func Q24Solution1(head *ListNode) *ListNode {
	p := head
	for p != nil {
		if p.Next != nil {
			//tmp := p.Val
			//p.Val = p.Next.Val
			//p.Next.Val=  tmp

			// 这里不用 tmp 变量来交换，按道理可以省掉一点内存，但从提交结果来看好像没有，可能是省的内存太少了，精度不够显示这部分内存。
			p.Val = p.Next.Val - p.Val
			p.Next.Val = p.Next.Val - p.Val
			p.Val = p.Val + p.Next.Val
		}

		if p.Next != nil {
			p = p.Next.Next
		} else {
			break
		}
	}

	return head
}