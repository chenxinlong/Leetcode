package algs

import "sort"

// 题287：https://leetcode.com/problems/find-the-duplicate-number/
// 找出数组中重复的元素(只有1个)

// 解法1：遍历并用 map 存储个数
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func Q287Solution1(nums []int) int {
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
		if m[n] > 1 {
			return n
		}
	}
	return 0
}

// 解法2：对 nums 排序，有序的 nums 如果有相邻值相等，那么它即重复元素
// 时间复杂度：等于 sort.Ints() 的复杂度
// 空间复杂度：O(1)
func Q287Solution2(nums []int) int {
	sort.Ints(nums)
	for i:=0; i<len(nums); i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return 0
}

// 解法3：环检测算法 (floyd's tortoise and hare)
// 记一个很有趣的视频：https://www.youtube.com/watch?v=pKO9UjSeLew&list=LLRW6F3yUZValAcvlaylffJA&index=2&t=0s
// todo
//func Q287Solution3(nums []int) int {
//
//}