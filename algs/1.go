package algs

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for k, v := range nums {
		if _, ok := m[v]; ok {
			return []int{m[v], k}
		}
		m[target-v] = k
	}

	return nil
}


/**
 * [analysis]
 * time complexity O(n)
 * use map
 */