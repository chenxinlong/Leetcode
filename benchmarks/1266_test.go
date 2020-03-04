package benchmarks

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
)
func BenchmarkQ1266(b *testing.B)  {
	b.Run("solution1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			algs.Q1266solution1(input_1266)
		}
	})
}