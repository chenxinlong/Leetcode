package algs

import (
	"math"
)

func TitleToNumber(s string) int {
	result := 0

	len := len(s)
	j := 0
	for i:=len-1; i >= 0; i-- {
		n := int(s[i] - 'A' + 1)
		result += n * int(math.Pow(26, float64(j)))
		j++
	}

	return result
}


/**
 * [analysis]
 * easy
 * Time complexity : O(n)
 */