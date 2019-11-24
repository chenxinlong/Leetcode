package algs

// 题5：https://leetcode.com/problems/longest-palindromic-substring/
// 求最长回文子串，很经典的一道题目，有各种解法，每种解法都值得学习。

// Brute force : 暴力法
// 按照惯例，一个解法献给 brute force。 获取一个s 的所有子串，判断是否回文，是则与上一个回文子串长度对比。
// 这的 2 个 for 嵌套就能得到 s 的所有子串。
// 时间复杂度 : O(n^3)
// 空间复杂度 : O(1)
func Q5Solution1(s string) string {
	sLen := len(s)
	if sLen < 2 {
		return s
	}

	// 判断 str 是否回文
	fnIsPalindromic := func(str string) bool {
		strLen := len(str)
		for left, right := 0, strLen-1; left < right; left, right = left+1, right-1 {
			if str[left] != str[right] {
				return false
			}
		}
		return true
	}

	result := ""
	// 遍历 s 的所有子串，逐个判断这个子串是否回文，如果子串回文并且长度大于上一个回文子串的长度，则把 result 替换为该子串。
	for i:=0; i<sLen; i++ {
		for j:=i+1; j<=sLen; j++ {
			if fnIsPalindromic(s[i:j]) && j-i >= len(result) {
				result = s[i:j]
			}
		}
	}

	return result
}
