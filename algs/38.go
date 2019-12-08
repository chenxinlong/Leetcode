package algs

import "strconv"

// 题38：https://leetcode.com/problems/count-and-say/

// 这题真的挺恶心的，看了讨论区，大部分人都没法理解 question description。而且理解之后发现题目本身更恶心。
func Q38Solution1(n int) string {
	if n == 1 {
		return "1"
	}

	result := handle([]byte("1"), n - 1)
	return string(result)
}

func handle(s []byte, n int) []byte {
	if n == 0 {
		return s
	}

	result, index, inputLength := []byte{}, 0, len(s)

	for index < inputLength {
		currentValue := s[index]
		lastRepeatedIndex := index

		for i := index; i < inputLength; i++ {
			if currentValue != s[i] {
				break
			}

			lastRepeatedIndex = i
		}

		length := lastRepeatedIndex - index + 1

		result = append(result, []byte(strconv.Itoa(length))...)
		result = append(result, currentValue)
		index = lastRepeatedIndex + 1
	}

	return handle(result, n-1)
}