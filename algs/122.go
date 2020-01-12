package algs

// 题122：https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

// 这道题其实如果理解透了其实比 121 题简单
// 题目要求多次买卖，得到最大收益。 开天眼而论，最大收益是什么？那就是只要后一天股价大于前一天，那么这部分涨幅就当做我们的收益。
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func Q122Solution1(prices []int) int {
	result := 0
	for i:=1; i<len(prices); i++ {
		if prices[i] > prices[i-1] {
			result += (prices[i] - prices[i-1])
		}
	}

	return result
}