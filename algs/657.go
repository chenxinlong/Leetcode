package algs

import "strings"

func judgeCircle(moves string) bool {
	return strings.Count(moves, "U") == strings.Count(moves, "D") && strings.Count(moves, "L") == strings.Count(moves, "R")
}

/**
 * [analysis]
 * If the num of two opposite direction are equal, then the vertical or horizontal coordinate will never change
 * Time complexity : O (n) if ignore strings.Count
 */
