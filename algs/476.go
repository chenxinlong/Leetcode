package algs

func FindComplement(num int) int {
	mask := 1
	tmpNum := num
	for tmpNum != 0 {
		tmpNum >>= 1
		mask <<= 1
	}

	return num ^ (mask - 1)
}


/**
 * [analysis]
 * Use the trick of Shift and XOR bitwise
 * For example : the complement num of 5 is [101] ^ [111], make mask as 1000, the result is [101] ^ ([1000] -1)
 */


 /**
  * [分析]
  * 使用移位运算符和异或运算符的性质
  * 比如 ： 5 的 complement num 是 [101] ^ [111] = 2, 只要用 n ^ [111....1] (n没有0前缀的 bit 数个 1, 5 的二进制是 101，所以 5^[111]) 即可
  */