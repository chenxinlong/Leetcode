package algs

// 题15：https://leetcode.com/problems/3sum/
// 这题的 2 个解法，都是看题解写出来的。 刚开始试了下暴力法(3层嵌套)，会产生一堆重复解，去重非常麻烦(除非每个解进行排序，然后对比)。

// 该题解是参考： https://leetcode.com/problems/3sum/discuss/110507/Golang-~n2%2Bn-worst-case-no-sort-no-deduplication-O(n2)-beats-50
// 这个解法的基础思路也是遍历每一个 i,j pair 的可能性，找到匹配的 k = -(i+j)，但是这里 i,j 不是通过双指针来找
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)，这里要衡量的不是 result，而是 m
func Q15Solution1(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}

	result := [][]int{}
	for i := range m {
		for j := range m {
			// i == j && m[i] == 1，说明此时 i 和 j 是同一位置的同一个元素，且这个元素只出现一次
			if i > j || (i==j && m[i] == 1) {
				continue
			}

			// i 和 j 是 pair, k 是补足数
			k := -(i+j)
			if k < i || k < j {
				continue
			}

			// 如果存在这个补足数
			if n, ok := m[k]; ok {
				if (k==i && n==1) || (k==j && n==1) || (i==j && j==k && n==2) {
					continue
				}
				result = append(result, []int{i, j, k})
			}
		}
	}

	return result
}

// 双指针解法
//func Q15Solution2(nums []int) [][]int {
//
//}