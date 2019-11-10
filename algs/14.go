package algs

// 题14：https://leetcode.com/problems/longest-common-prefix/

// 按照惯例，第一个解法都是首先想到的比较无脑的解法。 先把 strs 数组只有0个元素，1个元素这种 edge case 在第一步排除掉。
// 取 strs 第一个字符串作为基准字符串，剩下的字符串都和第一个字符串进行每个字符逐一对比。
// 当其中一个字符串的字符已经不够和第一个字符串的字符匹配了，或者任一字符串的第 j 个字符与第一个字符串的第 j 个字符不匹配，则此时就已经得到最后结果了。
// 时间复杂度：O(m*n),m 为 strs[0] 的长度，n 为 strs 数组的元素数量
// 空间复杂度：O(1)
// 这个解法和 leetcode 给出的题解 Vertical scanning 思路差不多
func Q14Solution1(strs []string) string {
	strsLen := len(strs)
	if strsLen < 1 {
		return ""
	}
	if strsLen == 1 {
		return strs[0]
	}

	result := ""
	l := len(strs[0])
	for j:=0; j<l; j++ {
		matchThisChar := true
		// 第一行和第一行一定完全匹配，所以跳过 i=0 的比较，i 直接从 1 开始
		for i:=1; i < strsLen; i++{
			// 如果当前字符串的长度已经不够与 strs[0] 匹配了，即不匹配，则退出
			if len(strs[i]) <= j {
				matchThisChar = false
				break
			}

			// 如果当前字符串[j] != 第一个字符串[j]，即不匹配，则退出
			if strs[i][j] != strs[0][j] {
				matchThisChar = false
				break
			}

		}

		if !matchThisChar {
			break
		} else {
			result += string(strs[0][j])
		}
	}

	return result
}

// 分治法 (Divide and conquer), 先分(divide) 后治(conquer)
//              {leetcode, leet, lee, le}
//         .-------------------------------------.
//  divide |         ↙                ↘         |
//         | {leetcode,leet}        {lee,le}     |
//         .-------------------------------------.
// conquer |        ↓                   ↓        |
//         |      {leet}               {le}      |
//         |         ↘                ↙         |
//         .-------------------------------------.
//                           {le}
// 我们把含有多个 string 的 strs 不断地从中间分为左右两半, 直到左半部分的数组和右边部分数组各自只有 1 个字符串，这个时候对比2个字符串的公共前缀就简单多了。
// 把规模n不断地分割成尽可能小的规模，直到我们可以很简单的计算这个小规模。 这题用递归方式分治，在每次递推过程中减少计算规模，当达到最小规模后进行计算，然后开始回归。
//
// 时间复杂度：O(m*n) m为 strs 里所有字符串的平均长度，n 为 strs 长度(这是因为这个 strs 里的每个字符串的每个字符都要参与判断)。 这个解法的核心代码在 fnCommPrefix 内部的公共前缀判断，因为一次 fnCommPrefix 的调用会对2个字符串进行判断，实际上应该是 (1/2)*m*n。
// 空间复杂度：O(log(n) * m) 递归调用，深度为 log(n)，即左右两半各进行 log(n) 次递归调用，会开辟 2log(n) 个栈空间。假设每个栈占 m 个空间，则共占用 2*log(n)*m 个空间。
func Q14Solution2(strs []string) string {
	// 按照惯例，先去掉讨厌的 edge case
	strsLen := len(strs)
	if strsLen < 1 {
		return ""
	}
	if strsLen == 1 {
		return strs[0]
	}

	// 获取字符串 a 和 b 的公共前缀
	fnCommPrefix := func(a, b string) string {
		minLen := len(a)
		if len(b) < minLen {
			minLen = len(b)
		}
		for i:=0; i<minLen; i++ {
			if a[i] != b[i] {
				return a[:i]
			}
		}
		return a[:minLen]
	}

	// 如果用 := 符声明匿名函数并实现函数体，那么该匿名函数无法递归调用。
	// 应该用下面的"先声明类型，再实现函数体" 的方式
	//fnLongestCommonPrefix := func(strs []string, left, right int) string {
	//	//	if left == right {
	//	//		return strs[left]
	//	//	}
	//	//
	//	//	mid := (left+right)/2
	//	//	lcpLeft := fnLongestCommonPrefix(strs, left, mid)
	//	//	lcpRight := fnLongestCommonPrefix(strs, mid+1, right)
	//	//	return fnCommPrefix(lcpLeft, lcpRight)
	//	//}
	var fnLongestCommonPrefix func(strs []string, left, right int) string
	fnLongestCommonPrefix = func(strs []string, left, right int) string {
		if left == right {
			return strs[left]
		}

		mid := (left+right)/2
		lcpLeft := fnLongestCommonPrefix(strs, left, mid)
		lcpRight := fnLongestCommonPrefix(strs, mid+1, right)
		return fnCommPrefix(lcpLeft, lcpRight)
	}
	return fnLongestCommonPrefix(strs, 0, len(strs)-1)
}


