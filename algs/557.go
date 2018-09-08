package algs

import (
	"strings"
)

func ReverseWords(s string) string {
	words := strings.Split(s, " ")
	var wordsReversed []string
	for _, word := range words {
		wordsReversed = append(wordsReversed, reverseString(word))
	}

	result := strings.Join(wordsReversed, " ")
	return result
}

func reverseString(s string) string {
	if s != "" {
		return reverseString(s[1:]) + s[:1]
	}

	return ""
}


/**
 * [analysis]
 * Nothing to say
 */