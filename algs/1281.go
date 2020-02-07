package algs

// 题1281：https://leetcode.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/

// 简单计算题
// 因为做乘法，所以这里的 product 初始值要为 1。
// 时间复杂度：O(log(n))
// 空间复杂度：O(1)
func Q1281Solution1(n int) int {
	product := 1
	sum := 0
	for ; n != 0; n/=10 {
		digit := n % 10
		product *= digit
		sum += digit
	}

	return product-sum
}