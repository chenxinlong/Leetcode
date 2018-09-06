package algs

func SelfDividingNumbers(left int, right int) []int {
	var nums []int

	for left <= right {
		tmpLeft := left
		isSelfDivding := true
		for tmpLeft != 0 {
			mod := tmpLeft % 10
			if mod == 0 {
				isSelfDivding = false
				break
			}

			if left % mod != 0 {
				isSelfDivding = false
				break
			}

			tmpLeft /= 10
		}

		if isSelfDivding == true {
			nums = append(nums, left)
		}

		left ++
	}

	return nums
}

/**
 * [analysis]
 * Nothing to say
 * Time complexity : O(n log 10n)
 */