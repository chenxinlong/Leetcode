package algs

// 题15：https://leetcode.com/problems/3sum/
// 这题的 2 个解法，都是看题解写出来的。 刚开始试了下暴力法(3层嵌套)，会产生一堆重复解，去重非常麻烦(除非每个解进行排序，然后对比)。

// 第一个题解是参考： https://leetcode.com/problems/3sum/discuss/110507/Golang-~n2%2Bn-worst-case-no-sort-no-deduplication-O(n2)-beats-50
// 这个解法的基础思路也是固定 k，找到匹配的 i+j 以满足 k = -(i+j)，但是 i,j 不是通过双指针来找
//func Q15Solution1(nums []int) [][]int {
//
//}