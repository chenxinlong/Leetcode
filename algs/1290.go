package algs

import (
	"math"
)

// 题1290：https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/

// 链表遍历 + 进制转换
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func S1290Solution1(head *ListNode) int {
	var result float64
	binaryIntArr := []int{}
	for head != nil {
		binaryIntArr = append(binaryIntArr, head.Val)

		head = head.Next
	}
	binaryLen := len(binaryIntArr)
	for i,j :=binaryLen-1, 0; i>=0; i,j = i-1,j+1 {
		result += float64(binaryIntArr[i]) * math.Pow(2, float64(j))
	}
	return int(result)
}

// 位运算
// 刚开始没想到这个解法，看了这道题的 related topics 有 "bit manipulation" 之后，去看了下社区讨论，发现其实这个解法最简单，只是没想到。
// 因为是从头结点遍历，所以访问到的每一个节点 result 都左移1位，然后位或("|") 这个节点的 val
// 这样可以直接构造出十进制的结果,也无需像 solutin 1 那样做各种类型转换.
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func S1290Solution2(head *ListNode) int {
	result := 0
	for head != nil {
		result = (result << 1) | head.Val
		head = head.Next
	}

	return result
}