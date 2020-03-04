package algs

import "math"

// 题1266：https://leetcode.com/problems/minimum-time-visiting-all-points/

// 解法1
// 以后如果有移动坐标的步数的计算，可以用这个规则计算 ：
// 1. 如果只能中规中矩横向/纵向移动一个单位，那么 p1 -> p2 所需移动最少步数是 math.abs(p1[x]-p2[x]) + math.abs(p1[y]-p2[y])
// 2. 如果像这题一样，可以向对角移动一个单位，那么 p1 -> p2 所需移动最少步数是 math.Max(math.abs(p1[x]-p2[x]), math.abs(p1[y]-p2[y]))
// 时间复杂度：O(n)
// 空间复杂度：O(n),因为每次调用 minStepToNext() 都开辟了一个栈空间, 如果我把这个函数内容拿出来直接过程化执行，空间复杂度就是 O(1)
func Q1266solution1(points [][]int) int {
	pointsLen := len(points)
	if pointsLen <= 1 {
		return 0
	}

	// 计算2个坐标之间的最少步数
	// [1,1] ~ [3,4] ~ [-1,0]
	//   p1      p2       p3
	//
	// 可以移对角，那这个就好算了。
	// 拿 math.abs(p1[x] - p2[x]) 和 math.abs(p1[y]-p2[y]) 对比一下，较大者就是所需最小步数。
	// [1,1] ~ [3,4] ，最小步数是 math.Max(math.abs(1-3), math.abs(1-4)) = 3
	// [3,4] ~ [-1,0], 最小步数是 math.Max(math.abs(3+1), math.abs(4-0)) = 4
	minStepToNext := func(this, next []int) int {
		minStep := math.Max(math.Abs(float64(this[0]-next[0])), math.Abs(float64(this[1]-next[1])))
		return int(minStep)
	}

	totalSteps := 0
	for i := 0; i < pointsLen-1; i++ {
		this, next := points[i], points[i+1]
		totalSteps += minStepToNext(this, next)
	}

	return totalSteps
}