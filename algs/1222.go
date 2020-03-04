package algs

import (
	"math"
)

// 题1222：https://leetcode.com/problems/queens-that-can-attack-the-king/
// 这题其实就是求与 king 在同一横、纵、对角的最近的 queen 的坐标
//        \|/
//       --k--
//        /|\
// 所以这道题最多返回 8 个 queen

// 解法1 ： 规规矩矩判断就好 。
// 存储这8个方位的 queen，每次发现距离king更近的queen，直接替换掉上一个距离 king 最近的 queen 就好。
// 总结，这个坐标系真的要写吐了
// 讨论：https://leetcode.com/problems/queens-that-can-attack-the-king/discuss/531751/Golang-easy-understand-golang-solution-with-linear-time-complexity
// 时间复杂度：O(n) n 是 queen 数量
// 空间复杂度：O(1) 相当于一个 8x2 矩阵，result 最多也就为 8。
func Q1222Solution1(queens [][]int, king []int) [][]int {
	kingX, kingY := king[0], king[1]
	nearestQueens := map[string][]int{
		"up":[]int{kingX, 8},
		"down":[]int{kingX, -8},
		"left":[]int{-8, kingY},
		"right":[]int{8,kingY},
		"up_left":[]int{-8, 8},
		"up_right":[]int{8, 8},
		"down_left":[]int{-8, -8},
		"down_right":[]int{8, -8},
	}
	for _, queen := range queens {
		queenX, queenY := queen[0], queen[1]
		// 在同一纵轴
		if queenX == kingX {
			if (queenY >= kingY) && (queenY < nearestQueens["up"][1]) {
				nearestQueens["up"] = queen
			}
			if (queenY <= kingY) && (queenY > nearestQueens["down"][1]) {
				nearestQueens["down"] = queen
			}
		}
		// 在同一横轴
		if queenY == kingY {
			if (queenX >= kingX) && (queenX < nearestQueens["right"][0]) {
				nearestQueens["right"] = queen
			}
			if (queenX <= kingX) && (queenX > nearestQueens["left"][0]) {
				nearestQueens["left"] = queen
			}
		}
		// 在同一对角线
		// 在这里吃了个小亏，刚开始是写 ：
		// if (queenX-kingX) == (queenY - kingY)
		// 发现漏了刚好一正一负的情况。
		if math.Abs(float64(queenX-kingX)) == math.Abs(float64(queenY-kingY)) {
			// 在同一对角线的 queen, 只要 x或y的任意值比 nearestQueens 里的同方向的 queen 更接近，那就是更接近
			if (queenX >= kingX) && (queenY > kingY) && (queenX < nearestQueens["up_right"][0]) {
				nearestQueens["up_right"] = queen
			}
			if (queenX >= kingX) && (queenY < kingY) && (queenX < nearestQueens["down_right"][0]) {
				nearestQueens["down_right"] = queen
			}
			if (queenX <= kingX) && (queenY > kingY) && (queenX > nearestQueens["up_left"][0]) {
				nearestQueens["up_left"] = queen
			}
			if (queenX <= kingX) && (queenY < kingY) && (queenX > nearestQueens["down_left"][0]) {
				nearestQueens["down_left"] = queen
			}
		}
	}


	// 把可以攻击 king 的 queen 放入结果集返回
	result := [][]int{}
	for _, queen := range nearestQueens {
		if queen[0] == 8 || queen[0] == -8 || queen[1] == 8 || queen[1] == -8 {
			continue
		}
		result = append(result, queen)
	}
	return result
}