package algs

import (
	"strings"
)

func NumJewelsInStones(J string, S string) int {
	var counts int
	for _, letter := range J {
		counts += strings.Count(S, string(letter))
	}

	return counts
}

/**
 *  [analysis]
 *  Take a brief look of discussion, seems there's no better way.
 *  Time complexity : O(n) if ignore strings.Count 
 */