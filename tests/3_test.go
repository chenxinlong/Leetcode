package tests

import (
	"github.com/chenxinlong/leetcode/algs"
	"testing"
)

func Test3(t *testing.T)  {
	cases := map[string]int {
		"abcabcbb":3,
		"bbbbb":1,
		"pwwkew":3,
		" ":1,
		"a":1,
		"abcdefg":7,
	}
	
	funcs := algs.Q3()
	
	for solutionName, fn := range funcs {
		for input, expected := range cases {
			if output := fn(input); output != expected {
				t.Errorf("Solution : [%v] Input : %v, Expected : %v, Output : %v", solutionName, input, expected, output)
			}
		}
	}
}