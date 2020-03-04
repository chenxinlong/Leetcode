package benchmarks

import (
	"github.com/chenxinlong/leetcode/algs"
	"testing"
)

var (
	input_1365 = []int{8,1,2,2,3}
)

func BenchmarkQ1365(b *testing.B)  {
	b.Run("solution1", func(b *testing.B) {
		for i := 0; i<b.N;i++  {
			algs.Q1365Solution1(input_1365)
		}
	})
}