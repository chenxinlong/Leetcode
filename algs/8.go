package algs

import "strings"

// 题8：https://leetcode.com/problems/string-to-integer-atoi/
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func Q8Solution1(str string) int {
	// 第一步 trim 掉头尾空格，这样不但可以在后续逻辑中省掉不少针对极端用例的判断, 还能让恶心的 +, - 判断变得简单点
	// 虽然用了 built-in api，但是如果不这样处理，这题后面就会多不少判断
	str = strings.TrimSpace(str)
	l := len(str)
	if l < 1 {
		return 0
	}

	result := 0
	const MaxInt32 = 1 << 31-1
	const MinInt32 = -1 << 31
	m := map[uint8]int{
		48:0,
		49:1,
		50:2,
		51:3,
		52:4,
		53:5,
		54:6,
		55:7,
		56:8,
		57:9,
	}

	// 我们假设合法的字符串应该第一位就可以判断出正负
	sign := 1
	for i:=0; i < l; i++ {
		// 第一次迭代(即 i==0 时), 就要得出是结果是正还是负。
		// 在 str 被 trim space 之后，再经过这一步后面就好判断了：+,-符只允许出现在 str 首位, 除首位以外的其余位如果不是数字就可以判定无法转为 int, 直接返回 0
		if i == 0 {
			if str[0] == 45 {
				sign = -1
				continue
			}
			if str[0] == 43 {
				continue
			}
		}

		// 非字符串首位的字符，只允许是数字
		if str[i] < 48 || str[i] > 57 {
			break
		}

		result = result * 10 + m[str[i]]

		// 上/下溢出返回32位整数上/下限
		if sign * result > MaxInt32  {
			return MaxInt32
		}
		if sign * result < MinInt32 {
			return MinInt32
		}
	}

	return  sign * result
}