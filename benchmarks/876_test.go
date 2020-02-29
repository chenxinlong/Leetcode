package benchmarks

import (
	"github.com/chenxinlong/leetcode/algs"
	"testing"
)

var (
	input_876 = algs.ArrayToSingleLinkedList([]int{1,2,3,4,5,6,7,8,9,10})
)

func BenchmarkQ876Solution1(b *testing.B) {
	for i := 0; i < b.N ; i++  {
		algs.Q876Solution1(input_876)
	}
}

func BenchmarkQ876Solution2(b *testing.B) {
	for i := 0; i < b.N ; i++  {
		algs.Q876Solution2(input_876)
	}
}