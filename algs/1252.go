package algs

// 题1252：https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/

// 解法1 ：第一个解法没啥好说的，直接规规矩矩一个个去算，一个个去数
// 时间复杂度：O((m + n) * k) 这里的 k = len(indices)
// 空间复杂度：O(m * n)
func Q1252Solution1(n int, m int, indices [][]int) int {
	// Go 里定义数组时，长度必须是 constant representable 的
	// 这里不能直接把 m,n 直接传进去，否则会报错：non-constant array bound n
	// 题目限制了 n 和 m 是 1~ 50，那我们直接造一个 50x50 的二维数组即可
	// matrix := [n][m]int{}

	matrix := [50][50]int{}
	indicesLen := len(indices)
	for i := 0; i < indicesLen; i++ {
		ri, ci := indices[i][0], indices[i][1]
		for j:=0;j<m;j++ {
			matrix[ri][j] ^= 1
		}
		for k :=0; k<n; k++ {
			matrix[k][ci] ^= 1
		}
	}

	result := 0
	for _, rows := range matrix {
		for _, num := range rows {
			if num  == 1 {
				result++
			}
		}
	}
	return result
}