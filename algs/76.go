package algs

// 题76：https://leetcode.com/problems/minimum-window-substring/

// 按照惯例，第一个解法献给暴力法
// 现在做题都有一个习惯，为了保证你的 main idea 落实得干净纯粹，就先把一些 corner case 在算法主流程开始之前，就全部去掉。避免在算法主流程中，加一堆关于这些边界值的判断，影响算法主要逻辑部分的美观。
// 这题和第 3 题的暴力解法差不多，遍历所有子串，再逐个判断。
// 时间复杂度：O(n^3)
// 空间复杂度：O(n)
func Q76Solution1(s string, t string) string {
	tLen := len(t)
	sLen := len(s)
	// 率先去掉 corner case，让后面的算法更纯粹和干净。
	if tLen < 1 {
		return ""
	}
	if sLen < tLen {
		return ""
	}

	// 判断字符串 sub 中是否包含字符串 t 所有字符, 包括出现次数
	includeAllT := func(subs string, t string) bool {
		if len(subs) < len(t) {
			return false
		}

		mt := map[int32]int{}
		ms := map[int32]int{}
		for _, j := range t {
			mt[j]++
		}

		for _, i := range subs {
			ms[i]++
		}

		for k,v := range mt {
			if ms[k] != v {
				return false
			}
		}

		return true
	}

	minStr := s
	leastIncludeOne := false
	for i:=0; i<sLen; i++ {
		for j:=i; j<sLen; j++ {
			if includeAllT(s[i:j+1], t) {
				leastIncludeOne = true
				if (j-i+1) < len(minStr) {
					minStr = s[i:j+1]
				}
			}
		}
	}

	if leastIncludeOne == false {
		return ""
	}

	return minStr
}