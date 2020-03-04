package algs

import "sort"

// 题128：https://leetcode.com/problems/longest-consecutive-sequence/
// 这道题排序可解，暴力法可解。只是不符合题中对时间复杂度 O(n) 的要求。

// 解法1 ： hash map
// 注：会 TLE 和 OOM  ，所以这个解法不用了
// 时间复杂度：O(贼大)
// 空间复杂度：O(贼大)
func Q128Soution1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	m := map[int]int{}
	min, max := nums[0], nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		m[n]++
	}


	length := 0
	lengths := []int{}
	last := min
	// 这种方式存在弊端，就是当  [-999999999,1,99999999999] 这种情况时，要把 -999999999 ~ 99999999999 中间的所有数字遍历一遍
	for i:=min; i <=max; i++ {
		_, ok := m[i];
		if i == min || (ok && i == last+1) {
			length++
			last = i
		} else {
			lengths = append(lengths, length)
			length = 0
			last = i
		}

		// 当 i == max 时，把最后的 length 塞进 lengths
		if i == max {
			lengths = append(lengths, length)
		}
	}

	result := 0
	for _, length := range lengths {
		if length > result {
			result = length
		}
	}
	return result
}


// 解法2：迫不得已用上了排序, 什么 O(n) 的要求，等优化轮再说吧。
// 时间复杂度：O(排序所需时间复杂度)
// 空间复杂度：O(排序所需空间复杂度)
func Q128Soution2(nums []int) int {
	numsLen := len(nums)
	if numsLen <= 1 {
		return numsLen
	}

	sort.Ints(nums)
	maxLength, currentLength := 0, 0
	for i:=0; i<numsLen; i++ {
		if i == 0 {
			currentLength++
			continue
		}
		if nums[i] == (nums[i-1]+1) {
			currentLength++
		} else if nums[i] == nums[i-1] {
			continue
		} else {
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 1
		}
	}

	if currentLength > maxLength {
		maxLength = currentLength
	}
	return maxLength
}