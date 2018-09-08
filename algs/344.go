package algs

// solution 1 : time limit exceed
func ReverseString1(s string) string {
	if s != "" {
		return ReverseString1(s[1:]) + s[:1]
	}

	return ""
}

// solution 2 : accepted
func ReverseString2(s string) string {
	var result string

	len := len(s)
	for ;len > 0;len-- {
		result += string(s[len-1])

	}
	return result
}


/**
 * [analysis]
 * solution 1 use recursion but can not pass all test case for the reason of time limit exceed
 */