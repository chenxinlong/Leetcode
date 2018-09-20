package algs

func Intersection(nums1 []int, nums2 []int) []int {
	numMap := map[int]bool{}
	for _, i := range nums1 {
		numMap[i] = true
	}

	intersection := []int{}
	for _,i := range nums2 {
		if numMap[i] == true {
			intersection = append(intersection, i)
			numMap[i] = false
		}
	}
	return intersection
}

/**
 * [analysis]
 * easy
 * Time complexity : O(n)
 */