package algs

func Q9() map[string]func(x int) bool {
	solutionFuncs := map[string]func(x int) bool{
		"1.Common": Q9Solution1,
		"2.Common optmized": Q9Solution2,
	}
	return solutionFuncs
}

// 题9：https://leetcode.com/problems/palindrome-number/

// 把 x 的每一位取出来，放到一个 int slice 中。头尾双指针遍历 slice ，每次迭代判断该 slice 头尾元素是否相等。
// 这样做增加了空间复杂度。
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func Q9Solution1(x int) bool {
	if x < 0 {
		return false
	}

	i := []int{}
	for ;x != 0; x /=10 {
		i = append(i, x%10)
	}

	for left, right := 0, len(i)-1; left < right; left, right = left+1, right-1 {
		if i[left] != i[right] {
			return false
		}
	}

	return true
}


// 把 x 不断除以10取余直到 x 为 0, 然后把每次得到的余数乘以10。相当于不断用x的余数重塑出一个 x 的逆数 y。如果结果 x 等于 y，那么 x 的字符串回文。
// 时间复杂度：O(log10(n))
// 空间复杂度：O(1)
func Q9Solution2(x int) bool {
	if x < 0 {
		return false
	}

	y := 0
	for tmpX := x; tmpX != 0; tmpX /=10 {
		y = y*10 + (tmpX%10)
	}

	return x == y
}