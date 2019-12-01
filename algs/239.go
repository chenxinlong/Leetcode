package algs

// 题239：https://leetcode.com/problems/sliding-window-maximum/

// 时间复杂度：O(n)
// 空间复杂度：O(1)
func Q239Solution1(nums []int, k int) []int {
	result := make([]int, 0, len(nums)-k)
	window := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if len(window) > 0 && window[0] <= i-k {
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] < nums[i] {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		if i >= k-1 {
			result = append(result, nums[window[0]])
		}
	}
	return result
}