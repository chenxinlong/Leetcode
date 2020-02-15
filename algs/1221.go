package algs

// 题1221：https://leetcode.com/problems/split-a-string-in-balanced-strings/

// brute force(暴力法)
// 时间复杂度：O(n^3)
// 空间复杂度：O(1) 这里的 m 实际上只存储 m["L"] 和 m["R"]，并且 m 的生命周期在每个子串的 balance 判断完后就结束了。
func Q1221Solution1(s string) int {
	isBalanced := func(s string) bool {
		sLen := len(s)
		m := map[string]int{}
		for i := 0; i<sLen; i++ {
			m[string(s[i])]++
		}
		return m["L"] == m["R"]
	}

	sLen := len(s)
	result := 0
	for i:=0; i<sLen; i++ {
		for j:=i; j<sLen; j++ {
			if isBalanced(s[i:j+1]) {
				result++
				// 其实题目不是判断所有子串
				// 如果是判断所有子串的话， LRLR 应该有 4 个 balance 子串 : LR, RL, LR, LRLR
				// 但是实际上只有 LR, LR 2 个，说明第一个 LR 判断是 balance 之后，直接跳过这个子串。
				i = j+1
			}
		}
	}

	return result
}