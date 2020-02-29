package benchmarks

import (
	"github.com/chenxinlong/leetcode/algs"
	"testing"
)

var (
	input_876 = algs.ArrayToSingleLinkedList([]int{1,2,3,4,5,6,7,8,9,10})
)

func BenchmarkQ876(b *testing.B) {
	b.Run("solution1", func(b *testing.B) {
		for i:=0; i<b.N; i++ {
			algs.Q876Solution1(input_876)
		}
	})

	b.Run("solution2", func(b *testing.B) {
		for i:=0; i<b.N; i++ {
			algs.Q876Solution2(input_876)
		}
	})
}