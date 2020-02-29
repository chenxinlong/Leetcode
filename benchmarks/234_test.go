package benchmarks

import (
	"github.com/chenxinlong/leetcode/algs"
	"testing"
)

var (
	input_234 = algs.ArrayToSingleLinkedList([]int{1,2,3,4,5,4,3,2,1})
)

func BenchmarkQ234(b *testing.B) {
	b.Run("solution1", func(b *testing.B) {
		for i:=0; i<b.N; i++ {
			algs.Q234Solution1(input_234)
		}
	})

	b.Run("solution2", func(b *testing.B) {
		for i:=0; i<b.N; i++ {
			algs.Q234Solution2(input_234)
		}
	})
}