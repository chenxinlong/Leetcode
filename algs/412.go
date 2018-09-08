package algs

import (
	"strconv"
)

func FizzBuzz(n int) []string {
	var result []string

	for i:=1; i<=n; i++ {
		switch {
		case i % 15 == 0:
			result = append(result, "FizzBuzz")
		case i % 5 == 0:
			result = append(result, "Buzz")
		case i % 3 == 0:
			result = append(result, "Fizz")
		default:
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}

/*
 * [analysis]
 * easy
 * Time complexity : O(n)
 */