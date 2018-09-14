package algs

func AddDigits(num int) int {

	for num >= 10 {
		n := num % 10
		num /= 10
		num += n
	}

	return num
}

/**
 * [analysis]
 * Question expect there's no loop/recursion in solution and with O(1) runtime, so there must be a math formula to deal with this question
 * Time complexity : O(n)
 */