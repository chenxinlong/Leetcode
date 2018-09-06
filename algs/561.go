package algs

import "sort"

func ArrayPairSum(nums []int) int {
	sort.Ints(nums)

	len := len(nums)
	sum := 0
	for i:=0; i < len; i+=2 {
		sum += nums[i]
	}

	return sum
}


/**
 * [analysis]
 * Sort this array, add the first element of each pair
 */