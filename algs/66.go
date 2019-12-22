package algs

// 题 66 ：https://leetcode.com/problems/plus-one/

// 这一题好像没什么难度。也没什么需要仔细注意的地方。就是.. 很简单。 好像也没什么特别 trick 的解法。
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func Q66Solution1(digits []int) []int {
	digitsLen := len(digits)
	if digitsLen < 1 {
		return digits
	}

	result := []int{}
	carry := 1
	for i :=digitsLen-1; i>=0; i-- {
		sum := digits[i]+carry
		digit := sum%10
		if sum < 10 {
			carry = 0
		}

		result = append([]int{digit}, result...)
	}
	if carry != 0 {
		result = append([]int{carry}, result...)
	}

	return result
}