package algs

func PeakIndexInMountainArray(A []int) int {
	i := 0
	for A[i] < A[i+1]  {
		i++
	}

	return i
}

/**
 * [analysis]
 * easy for linear scan with O(n) time complexity
 */
