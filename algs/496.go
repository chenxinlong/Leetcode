package algs

func NextGreaterElement(findNums []int, nums []int) []int {
	var result []int
	for _, n := range findNums {
		r := -1
		nPassed := false
		for _, m := range nums {
			if m == n {
				nPassed = true
			}
			if m > n && nPassed {
				r = m
				break
			}
		}

		result = append(result, r)
	}

	return result
}

/**
 * [analysis]
 * easy
 * Time complexity : O(n^2)
 */