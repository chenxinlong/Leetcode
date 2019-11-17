package algs

// 题26：https://leetcode.com/problems/remove-duplicates-from-sorted-array

// 这题不仅要返回最后剩余的 nums 数量就行，而是要把重复数删除。
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func Q26Solution1(nums []int) int {
	i := 0
	for j:=1; j < len(nums); j++{
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i+1
}