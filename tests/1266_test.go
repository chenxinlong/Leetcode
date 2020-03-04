package tests

import (
	"github.com/chenxinlong/leetcode/algs"
	"testing"
)

var (
	input_1266 = [][]int{
		{1,1},
		{3,4},
		{-1,0},
	}
	output_1266 = 7
)

func Test1266(t *testing.T) {
	t.Run("solution1", func(t *testing.T) {
		want := output_1266
		got := algs.Q1266solution1(input_1266)
		if got != want {
			t.Errorf("want : %+v, got : %+v", want, got)
		}
	})
}