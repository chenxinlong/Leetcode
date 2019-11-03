package algs

// 题7：https://leetcode.com/problems/reverse-integer/
// 32位有符号整数，上下溢出则返回0
// -123 % 10 = -3
// 时间复杂度：O(log(n))
// 空间复杂度：O(1)
func Q7Solution1(x int) int {
	const MaxInt32 = 1 << 31-1
	const MinInt32 = -1 << 31
	var res int

	for ;x != 0; x/=10 {
		res = res*10 + x%10

		// 每次累加之后都判断当前 res 是否上/下溢，一旦溢出，就直接返回0
		if res < MinInt32 || res > MaxInt32 {
			return 0
		}
	}

	return res
}