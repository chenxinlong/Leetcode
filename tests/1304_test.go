package tests

import (
	"github.com/chenxinlong/leetcode/algs"
	"reflect"
	"testing"
)

// 写单元测试
// 给单元测试写单元测试
// 给单元测试的单元测试写单元测试
// ...
func Test1304(t *testing.T)  {
	t.Run("solution1_case1", func(t *testing.T) {
		got := algs.Q1304Solution1(1)
		want := []int{0}
		if reflect.DeepEqual(got, want) == false {
			t.Errorf("want : %v, got : %v", want, got)
		}
	})

	t.Run("solution1_case_all", func(t *testing.T) {
		n := 10
		got := algs.Q1304Solution1(n)
		if len(got) != n {
			t.Errorf("length incorrect as : %v", len(got))
			return
		}

		uniqueMap := map[int]int{}
		for _, i := range got {
			uniqueMap[i]+=1
			if uniqueMap[i] > 1 {
				t.Errorf("invalid due to duplicate element : %v", uniqueMap[i])
			}
		}
	})
}