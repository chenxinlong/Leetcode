package benchmarks

import (
	"github.com/chenxinlong/leetcode/algs"
	"testing"
)

var (
	input_287 = []int{1,3,2,5,4,6,9,8,11,23,12,42,23,1}
)

func BenchmarkQ287(b *testing.B)  {
	b.Run("solution1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			algs.Q287Solution1(input_287)
		}
	})

	b.Run("solution2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			algs.Q287Solution2(input_287)
		}
	})
}