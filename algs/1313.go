package algs

// 题1313：https://leetcode.com/problems/decompress-run-length-encoded-list/

// 这题没啥做的意思，差评。
// 以这个解法来说，注意防止数组越界即可。
// 时间复杂度：O(n* avg(freq)) 这题时间复杂度不好分析。遍历 nums 是 n/2，但 freq 无法确定。所以这里暂时用 avg(freq) 来代替吧。
// 空间复杂度：O(n* avg(freq)
func Q1313Solution1(nums []int) []int {
	result := []int{}
	numsLen := len(nums)
	for i:=0; i<numsLen-1; i+=2 {
		freq, val := nums[i], nums[i+1]
		for j:=0; j<freq; j++ {
			result = append(result, val)
		}
	}
	return result
}