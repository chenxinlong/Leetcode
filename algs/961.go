package algs

func RepeatedNTimes(A []int) int {
	mapN := map[int]bool{}
	for _, n := range A {
		if _, ok := mapN[n]; ok {
			return n
		}

		mapN[n] = true
	}

	return A[0]
}


/**
 * [analysis]
 * easy
 */