package algs

// 题1351：https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/

// 解法1：直接遍历不解释
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
// 我看了一下耗时排在我前面的解法，无脑遍历但是中间不 break 的，居然比我更快。。想不通。
func Q1351Solution1(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] < 0 {
				result += (n-j)
				break
			}
		}
	}

	return result
}