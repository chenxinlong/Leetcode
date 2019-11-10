package algs

// 题13：https://leetcode.com/problems/roman-to-integer/submissions/

// 这题和 12题相反，首先想到的解法也是这种很简单的，没啥好分析的
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func Q13Solution1(s string) int {
	m := map[string]int{"I":1, "IV":4, "V":5, "IX":9, "X":10, "XL":40, "L":50, "XC":90, "C":100, "CD":400, "D":500, "CM":900, "M":1000,}
	result := 0
	l := len(s)
	for i:=0; i < l; i++ {
		if i+1 < l  {
			if n, ok := m[string(s[i])+string(s[i+1])]; ok {
				result += n
				i++
				continue
			}
		}
		result += m[string(s[i])]
	}

	return result
}
