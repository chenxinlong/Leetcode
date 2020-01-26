package algs

import (
	"fmt"
	"strings"
)

// 题1108：https://leetcode.com/problems/defanging-an-ip-address/

// 这题就没啥好说的，应该说没有做的意义。
// 时间复杂度：O(1) ipv4 的格式固定，所以不论是 strings.split 还是下面的 for 循环的规模都是常数规模。
// 空间复杂度：O(1) ipv4 的格式固定，所以 IPParts 是一个常数复杂度。
func Q1108Solution1(address string) string {
	IPParts := strings.Split(address, ".")
	result := ""
	for i:=0; i<4; i++ {
		if i != 3 {
			result += fmt.Sprintf("%v[.]", IPParts[i])
		} else {
			result += IPParts[i]
		}
	}
	return result
}