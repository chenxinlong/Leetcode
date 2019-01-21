package algs

func SortArrayByParity(A []int) []int {
	leftN := 0
	for k, v := range A {
		if v % 2 == 0 && leftN <= k {
			A[k] = A[leftN]
			A[leftN] = v
			leftN ++
		}
	}

	return A
}

/**
 * [analysis]
 * leftN 表示左边为偶数的个数，把所有的偶数左移至 A[leftN], 原有 A[leftN] 交换至当前位置
 */