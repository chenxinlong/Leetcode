package algs

func FlipAndInvertImage(A [][]int) [][]int {
	for i, row := range A {
		if len(row) < 1 {
			continue
		}

		left := 0
		right := len(row)-1
		for left <= right {
			// 取反
			A[i][left] ^= 1
			// 如果 left == right, 连续取反2次又会变回原值
			if left != right {
				A[i][right] ^= 1
			}

			// 交换, 当 left == right, A[i][left] 和 A[i][right] 指向同一个值, 用加减法交换结果会变成 0
			if left != right {
				A[i][left] -= A[i][right]
				A[i][right] += A[i][left]
				A[i][left] = A[i][right] - A[i][left]
			}

			left++
			right--
		}
	}

	return A
}

/**
 * [analysis]
 * easy
 */