 package algs

 func MoveZeroes(nums []int)  {
	pre := 0
	for i, n := range nums {
		if n != 0 {
			nums[pre] = n

			// 如果位置变了，说明当前位置要变为 0
			if i != pre {
				nums[i] = 0
			}
			pre++
		}
	}
 }

 /**
  * [analysis]
  * Time complexity : O(n)
  */