package algs

import "fmt"

func MajorityElement(nums []int) int {
	mid := len(nums) / 2    // int type math.floor
	maps := map[int]int{}

	for _, n := range nums {
		if maps[n]++; maps[n] > mid {
			fmt.Println(maps[n])
			return n
		}
	}

	return nums[0]
}


/**
 * [analysis]
 * easy
 * Time complexity : O(n)
 */