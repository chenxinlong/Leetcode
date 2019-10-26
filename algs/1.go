package algs

func Q1() map[string] func(nums []int, target int) []int {
	solutionFuncs := map[string] func(nums []int, target int) []int {
		"1. Map":Q1Solution1,
	}

	return solutionFuncs
}

// 题1：https://leetcode.com/problems/two-sum/
// 遍历 nums ，边查(从map中查)边插(插入到map)
// "查" 指的是从 map 中查询期望匹配值，如果匹配(即相加为  target) 则返回结果
// "插" 指的是如果 map 中查不到期望匹配值时, 把自己的 value, key 也插入到 map 中  
// 时间复杂度 : O(n)   
// 空间复杂度 : O(n) 
func Q1Solution1(nums []int, target int) []int {
	m := map[int]int{}
	for k, v := range nums {
		if k2, ok := m[target-v]; ok {
			return []int{k2, k}
		}
		m[v] = k
	}

	return nil
}