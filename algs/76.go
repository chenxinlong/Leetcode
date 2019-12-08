package algs

// 题76：https://leetcode.com/problems/minimum-window-substring/

// Brute force : 按照惯例，第一个解法献给暴力法
// 现在做题都有一个习惯，为了保证你的 main idea 落实得干净纯粹，就先把一些 corner case 在算法主流程开始之前，就全部去掉。避免在算法主流程中有一堆关于这些边界值的判断，影响算法主要逻辑部分的可阅读性。
// 这题和第 3 题的暴力解法差不多，遍历所有子串，再逐个判断。实际上这个解法 TLE 了，所以直接前往下一个解法。
// 时间复杂度：O(n^3)
// 空间复杂度：O(n)
func Q76Solution1(s string, t string) string {
	tLen := len(t)
	sLen := len(s)
	// 率先去掉 corner case，让后面的算法更纯粹和干净。
	if sLen < tLen || s == "" || t == "" {
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
			if ms[k] < v {
				return false
			}
		}

		return true
	}

	minStr := s
	leastIncludeOne := false
	for i:=0; i<sLen; i++ {
		for j:=i; j<sLen; j++ {
			if j-i+1 > len(minStr) {
				break
			}
			if includeAllT(s[i:j+1], t) {
				leastIncludeOne = true
				if (j-i+1) < len(minStr) {
					minStr = s[i:j+1]
					break
				}
			}
		}
	}

	if leastIncludeOne == false {
		return ""
	}

	return minStr
}

// Slide window (滑动窗口)
// 窗口滑动并不一定是 i 和 j 同时自增或自减，也有可能是一次迭代只有 i 或者 j 增减，但只要 window 在整个流程中是往某个方向移动了，那么它就是滑动窗口。
// 这个解法的窗口还是右滑，只是迭代过程中的窗口大小变换规则和第 3 题不太一样，这题的窗口大小变换规则更复杂一点。
//
// 时间复杂度：O(n) n表示s长度。 注意虽然有嵌套循环，但是最差的情况是外层迭代一次，内层跟随着迭代一次, 所以最差情况也是 O(2n)。
// 空间复杂度：O(n+m) n 表示 s 唯一字符数，m 表示 t 唯一字符数。
func Q76Solution2(s string, t string) string {
	sLen := len(s)
	tLen := len(t)
	if sLen < tLen || s == "" || t == "" {
		return ""
	}

	mt := map[uint8]int{}
	ms := map[uint8]int{}
	for i:=0; i<tLen; i++ {
		mt[t[i]]++
	}

	// required 表示 t 中唯一字符个数
	// formed 表示当前 window 中已经满足了 t 中的字符的个数.
	// 比如 t = "abcc", s = "abbbbcc" ，那么 required 就是 3。假如当前 window 是 s[0]~s[5] = "abbbbc"，那么当前 window 已经满足了 t 中 "a" 和 "b" 的要求，但是 "c" 的数量还缺一个,s 所以此时 formed = 2
	var formed, required = 0, len(mt)
	var left, right = 0, 0

	// {窗口大小, left, right}
	var ans = []int{-1, 0, 0}

	// 窗口右滑，所以 right 触尾则循环结束
	for right < sLen {
		ms[s[right]]++
		if mt[s[right]] > 0 && mt[s[right]] == ms[s[right]] {
			formed++
		}

		// 当 formed 已经满足 required，则开始逐单元缩减左窗，以求得到一个既满足 formed == required 且最小的窗口
		for left <= right && formed == required {
			// 如果当前窗口小于 ans[0]，则该窗口是比 ans[0] 更接近答案的窗口，替换之.
			if (ans[0] == -1) || (right - left + 1 < ans[0]) {
				ans[0] = right - left + 1
				ans[1] = left
				ans[2] = right
			}

			ms[s[left]]--
			if mt[s[left]] > 0 && ms[s[left]] < mt[s[left]] {
				formed--
			}

			left++
		}

		// 如果当前窗口 formed 不满足 required，继续向右逐单元扩大窗口
		right++
	}

	if ans[0] == -1 {
		return ""
	}

	return s[ans[1] : ans[2]+1]
}
