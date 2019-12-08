package algs

// 题239：https://leetcode.com/problems/sliding-window-maximum/

// 第一个解法献给暴力法
// 时间复杂度：O(n * (k/2)) = O(n*k)
// 空间复杂度：O(1)
func Q239Solution1(nums []int, k int) []int {
	numsLen := len(nums)

	// 首先去掉 corner case，保证后面算法流程清晰
	if k == 1 || numsLen < 1{
		return nums
	}

	result := []int{}
	for i:=0; k<=numsLen; i,k=i+1,k+1 {
		windowLeft, windowRight, max := i, k-1, 0
		for windowLeft <= windowRight {
			if (nums[windowLeft] >= nums[windowRight]) && (nums[windowLeft] > max) {
				max = nums[windowLeft]
			}
			if (nums[windowRight] >= nums[windowLeft]) && (nums[windowRight] > max) {
				max = nums[windowRight]
			}
			windowLeft++
			windowRight--
		}

		result = append(result, max)
	}
	return result
}