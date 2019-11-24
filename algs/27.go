package algs

import "fmt"

// 题27：https://leetcode.com/problems/remove-element/solution/

// 这道题过了一周，还是不知道意义在哪里，先这样吧。。
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func Q27Solution1(nums []int, val int) int {
	i := 0
	numsLen := len(nums)
	for j:=0; j < numsLen; j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	fmt.Println(nums)
	return i
}