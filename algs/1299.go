package algs

// 题1299： https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/

// 解法1：这个解法是看题解才会做的，不然一开始想的也是暴力法。
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func Q1299Solution1(arr []int) []int {
	rightMax := -1
	for i := len(arr)-1; i >= 0; i-- {
		tmp := arr[i] // 防丢变量
		arr[i] = rightMax
		if tmp > rightMax {
			rightMax = tmp
		}
	}
	return arr
}