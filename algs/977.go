package algs

import "sort"

func SortedSquares(A []int) []int {
	res := []int{}
	for _, n := range A {
		res = append(res, n*n)
	}

	sort.Ints(res)

	return res
}


/**
 * [analysis]
 * sort
 */