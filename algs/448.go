package algs

func FindDisappearedNumbers(nums []int) []int {
	for i := range nums {
		for nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	result := []int{}
	for i, v := range nums {
		if v != i+1 {
			result = append(result, i+1)
		}
	}
	return result
}

/**
 * [analysis]
 * First nested loop convert nums to a sequence-correct array
 * Last loop retrieve the wrong element
 * Time complexity : O (n)
 */