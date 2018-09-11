package algs

func FindMaxConsecutiveOnes(nums []int) int {
	max := 0
	bit1 := 0
	for _, n := range nums {
		if n == 1 {
			bit1 ++
			if bit1 > max {
				max = bit1
			}
		}

		if n == 0 {
			bit1 = 0
		}
	}

	return max
}

/**
 * [analysis]
 * easy
 * time complexity : O(n)
 */