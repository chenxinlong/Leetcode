package benchmarks

import (
	"github.com/chenxinlong/leetcode/algs"
	"testing"
)

var (
	input_234 = algs.ArrayToSingleLinkedList([]int{1,2,3,4,5,4,3,2,1})
)

func BenchmarkQ234Solution1(b *testing.B)  {
	for i:=0; i<b.N; i++ {
		algs.Q234Solution1(input_234)
	}
}