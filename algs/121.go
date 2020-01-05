package algs

// 题121：https://leetcode.com/problems/best-time-to-buy-and-sell-stock/solution/

// 解法1：暴力法不解释
// 时间复杂度：O(n^2)
// 空间复杂度： O(1)
func Q121Solution1(prices []int) int {
	result := 0
	for i:=0; i < len(prices); i++ {
		for j:=i+1; j<len(prices);j++ {
			if prices[j] - prices[i] > result {
				result = prices[j] - prices[i]
			}
		}
	}
	return result
}