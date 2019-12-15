package main

import (
	"fmt"
	"github.com/chenxinlong/leetcode/algs"
)

func main () {
	a := &algs.ListNode{
		Val:1,
		Next:&algs.ListNode{
			Val:4,
			Next:&algs.ListNode{
				Val:5,
				Next:nil,
			},
		},
	}
	b := &algs.ListNode{
		Val:1,
		Next:&algs.ListNode{
			Val:3,
			Next:&algs.ListNode{
				Val:4,
				Next:nil,
			},
		},
	}
	c := &algs.ListNode{
		Val:2,
		Next:&algs.ListNode{
			Val:6,
			Next:nil,
		},
	}


	fmt.Printf("%+v", algs.Q23Solution2([]*algs.ListNode{a,b,c}))
}

