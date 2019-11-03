package algs

import (
	"math"
)

func Q3() map[string]func(s string) int {
	solutionFuncs := map[string]func(s string) int{
		"1.Brute force": Q3Solution1,
		"2.Slide window": Q3Solution2,
	}
	return solutionFuncs
}

// 题3 ：https://leetcode.com/problems/longest-substring-without-repeating-characters/

// Brute force(暴力解法)
// 时间复杂度 : O(n^3)
// 空间复杂度 : O(n)
// 遇到长文本 case 这个算法必然超时
func Q3Solution1(s string) int {
	l := len(s)
	var result float64 = 0

	// 判断 s 中的 start ~ end 这个子串中的字符是不是全部不重复
	isUnique := func(s string, start, end int) (result bool) {
		m := map[byte]bool{}
		for start <= end {
			if _, ok := m[s[start]]; ok {
				return false
			}
			m[s[start]] = true
			start++
		}

		return true
	}

	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if isUnique(s, i, j) {
				result = math.Max(result, float64(j-i+1))
			}
		}
	}

	return int(result)
}


// Slide window (滑动窗口)
// 这种滑动窗口的用法很经典，一定要掌握，常在字符串、数组这种线性结构中用到。
// 窗口可以是：
//    1. 初始状态i:0,j:0，窗口向右扩展范围、右滑
//    2. 初始状态i:0,j:(n-1)，窗口向左边缩小范围
//    3. ...
// 时间复杂度: 最好O(1), 平均O(n), 最差O(n) : 最好的情况:字符串小于2位数，平均情况=最差情况:除了最好情况外的其他所有情况
// 空间复杂度: 最好O(1), 平均O(n), 最差O(n) : 同上
func Q3Solution2(s string) int {
	l := len(s)
	if l < 2 {
		return l
	}

	result, i ,j := 0,0,0
	m := map[uint8]bool{}

	// [i, j) 表示这个 window
	// 当 window 的尾端 j 触尾之后说明 window 已经滑到最后了，就可以退出循环
	// 当 map 中无与 j 重复的值，该 window 的范围向右扩大一个单位
	// 当 map 中有与 j 重复的值，该 window 整体右滑一个单位
	for j < l {
		if _, ok := m[s[j]]; ok {
			// 窗口整体右滑 1 个单位,
			// 为什么窗口整体右滑 1 个单位是删除 s[i] 然后i++ ,而不是i++,j++ 呢？
			// 因为我们知道要整体右滑1个单位，所以 i++ 之后，原有的i已经不在窗口范围内了，那么原有的i不能对后面的结果产生影响，一定要从map中删除
			// 删除了原有的 i 之后，原有的 j 还没有进入到 map, 所以 j 这次不++，让 j 在下次的判断中进入到 map 中。
			delete(m , s[i])
			i++
		} else {
			m[s[j]] = true
			
			// 窗口范围向右扩大 1 个单位
			j++

			// 此时,范围扩大后的窗口如果比之前记录的最大窗口还要大，则取代原先结果
			if j-i > result {
				result = j-i
			}
		}
	}
	
	return result
}


