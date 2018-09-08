package algs

import (
	"strings"
)

func FindWords(words []string) []string {
	var result []string

	keyboard := map[int]string{1:"qwertyuiopQWERTYUIOP", 2:"asdfghjklASDFGHJKL", 3:"zxcvbnmZXCVBNM"}

	for _, word := range words {
		var r int
		for i, row := range keyboard {
			if strings.Contains(row, string(word[0])) {
				r = i
				break
			}
		}


		inSameRow := true
		for _, letter := range word {
			if strings.Contains(keyboard[r], string(letter)) == false {
				inSameRow = false
				break
			}
		}

		if inSameRow == true {
			result = append(result, word)
		}
	}

	return result
}

/**
 * [analysis]
 * It is easy but can't avoid nested loop
 * Time complexity : O(n^2)
 */