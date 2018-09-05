package algs

func HammingDistance(x int, y int) int {
	return countBits1(x ^ y)
}

func countBits1(n int) int {
	var counts int
	for n != 0 {
		if n%2 == 1 {
			counts++
		}
		n /= 2
	}

	return counts
}

/**
 * [analysis]
 * Use the properties of bitwise XOR, count the num of 1 in x ^ y
 * Time complexity : O(log n)
 */
