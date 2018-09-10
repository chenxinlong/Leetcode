package algs

func SingleNumber(nums []int) int {
	var result int

	for _, n := range nums {
		result ^= n
	}

	return result
}

/**
 * [analysis]
 * use the properties of bitwise XOR
 * time complexity : O(n)
 */