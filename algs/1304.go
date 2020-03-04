package algs

// 题1304：https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/
// 题目没看清，要求每个元素不同，所以一上来直接就给了一个含有 n 个0 的数组。

// 解法1 ： 比如 n = 5 , 直接构造一个 [1,2,3,4, -(1+2+3+4)] 即可。
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func Q1304Solution1(n int) []int {
	result := []int{}
	sum := 0
	for i := 1; i <= n-1; i++ {
		sum += i
		result = append(result, i)
	}
	result = append(result, -sum)
	return result
}