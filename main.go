package main

import (
	"fmt"
	"github.com/chenxinlong/leetcode/algs"
)

func main () {
	l := algs.ArrayToSingleLinkedList([]int{1,2,3,2,1})
	fmt.Println(algs.Q234Solution2(l))
}

