package tests

import (
	"github.com/chenxinlong/leetcode/algs"
	"testing"
)

type Q1Case struct {
	Nums []int 
	Target int
	Expected []int
}

func Test1(t *testing.T)  {
	cases := []Q1Case{
		{
			Nums:[]int{2,7,11,15},
			Target:9,
			Expected:[]int{0,1},
		},
	}

	funcs := algs.Q1()

	for solutionName, fn := range funcs {
		for _, testCase := range cases {
			output := fn(testCase.Nums, testCase.Target)
			for k, v := range output {
				if v != testCase.Expected[k] {
					t.Errorf("Solution : [%v] Input : %+v, Expected : %v, Output : %v", solutionName, testCase, testCase.Expected, output)
					break
				}
			}
		}
	}
}