package algs

// 题 1365:https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/

// 这是道简单题，只是这个解法我是看题解的。最开始我想的暴力法复杂度太高了，TLE 无疑。所以直接看题解。
// 时间复杂度：O(n)
// 空间复杂度：O(1)，因为题目说了最多 100 个元素，所以 m, n 数组不随 nums 规模而增大.
func Q1365Solution1(nums []int) []int {
	// 假设 input 是 []int{0,1,1,3,3,3,3,3,4}

	// 得到：
	// m[
	//   0 : 1
	//   1 : 2
	//   3 : 5
	//   4 : 1
	// ]
	// 即0有1个，1有2个，3有5个，4有1个
	m := [101]int{}
	for _, n := range nums {
		m[n]++
	}

	// 得到：
	// n [
	//     0 : 1
	//     1 : 3
	//     3 : 8
	//     4 : 9
	// ]
	// 即，>= 0 的有1个， >=1 的有3个， >=3 的有8个， >=4 的有9个。
	n := [101]int{}
	for k, v := range m {
		if k == 0 {
			n[k] = v
			continue
		}
		n[k] = v + n[k-1]
	}

	// n[0] 是>= 0 的个数，m[0] 是0的个数，所以 n[0]-m[0] 就是 > 0 的个数。
	result := []int{}
	for _, v := range nums {
		result = append(result, n[v]-m[v])
	}

	return result
}
