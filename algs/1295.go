package algs

// 题1295：https://leetcode.com/problems/find-numbers-with-even-number-of-digits/

// 简单题简单解
// 时间复杂度：O(n * log(n))  这里要注意，即使底数是 10，复杂度实际为 log10(n)，我们也认为它是log(n) 量级。
// 空间复杂度：O(1)
func Q1295Solution1(nums []int) int {
	isDigitLenEven := func(x int) bool {
		len := 0
		for x != 0 {
			x/=10
			len++
		}
		return len%2==0
	}

	result := 0
	for _, n := range nums {
		if isDigitLenEven(n) {
			result++
		}
	}

	return result
}