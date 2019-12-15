package algs

// 题 53：https://leetcode.com/problems/maximum-subarray/

// 按照惯例，第一个解法献给暴力法
// 这题无脑嵌套循环，就能解，但是毫无意外地会 TLE ，所以直接前往下一个题解。
// 时间复杂度：O(n^3)
// 空间复杂度：O(1)
func Q53Solution1(nums []int) int {
	numsLen := len(nums)
	if numsLen < 1 {
		return 0
	}

	// 这里踩了个小坑，max 不能直接初始化为 0，因为本题的 max 可能是负数。
	// 比如 nums = []int{-1,-2}，如果 max 初始化为0，那么这题就会返回 max = 0，而实际 max = -1
	// 所以直接初始化 max 为 nums 第一个元素即可。
	max := nums[0]

	for i:=0; i<numsLen; i++ {
		for j:=i; j<numsLen; j++ {
			subarray := nums[i:j+1]
			sum := 0
			for _, k := range subarray {
				sum += k
			}
			if sum > max {
				max = sum
			}
		}
	}

	return max
}

// 因为要求连续子数组，所以这题我们思路和第 3 题一样，用 slide window 来减少不必要的子数组比较。
func Q53Solution2(nums []int) int {
	numsLen := len(nums)
	if numsLen < 1 {
		return 0
	}

	max := nums[0]
	for i,j :=0,0; j<numsLen;  {
		subarray := nums[i:j+1]
		sum := 0
		for _, k := range subarray {
			sum += k
		}
		if sum >= max {
			max = sum
			j++
		} else {
			if i < j {
				i++
			} else {
				j++
			}
		}
	}

	return max
}

// 这个解法来自 ：https://leetcode.com/problems/maximum-subarray/discuss/119798/Greedy-Go-100-solution
func Q53Solution3(nums []int) int {
	max := -(1<<63)
	crt := 0

	for _, v := range nums {
		if crt + v < v {
			crt = v
		} else {
			crt += v
		}

		if crt > max {
			max = crt
		}
	}

	return max
}