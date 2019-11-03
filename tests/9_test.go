package tests

import (
	"github.com/chenxinlong/leetcode/algs"
	"testing"
)

func Test9(t *testing.T)  {
	cases := map[int]bool {
		-1:false,
		0:true,
		123:false,
		12321:true,
	}

	funcs := algs.Q9()

	for solutionName, fn := range funcs {
		for input, expected := range cases {
			if output := fn(input); output != expected {
				t.Errorf("Solution : [%v] Input : %v, Expected : %v, Output : %v", solutionName, input, expected, output)
			}
		}
	}
}