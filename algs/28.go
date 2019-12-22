package algs

// 题28：https://leetcode.com/problems/implement-strstr/
// 字符串匹配算法

// 按照惯例，第一个解法献给暴力法，逐字符匹配，没啥好说的。
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func Q28Solution(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	if haystackLen == needleLen && haystackLen == 0 {
		return 0
	}
	for i:=0; i<haystackLen; i++ {
		match := true
		tmpI := i
		for j := 0; j<needleLen; j++ {
			if tmpI >= haystackLen || haystack[tmpI] != needle[j] {
				match = false
				break
			} else {
				tmpI++
			}
		}

		if match == true {
			return i
		}
	}

	return -1
}

// 这个解法就是大名鼎鼎的 KMP 算法，实际上我到现在还没搞清楚 KMP 中失配函数和移位的关联是怎么证明出来的。
func Q28Solution2(haystack string, needle string) int {
	return -1
}