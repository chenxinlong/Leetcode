package algs

// 题12：https://leetcode.com/problems/integer-to-roman/submissions/

// 这种解法是最容易想到的，就没啥好分析的
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func Q12Solution1(num int) string {
	m := map[int]string{1:"I", 4:"IV", 5:"V", 9:"IX", 10:"X", 40:"XL", 50:"L", 90:"XC", 100:"C", 400:"CD", 500:"D", 900:"CM", 1000:"M",}
	arr := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}

	result := ""
	for _, d := range arr {
		for i:=0; i<num/d; i++ {
			result += m[d]
		}
		num %= d
		if num == 0 {
			break
		}
	}

	return result
}

// 这种解法虽然比较暴力，但又很聪明。 因为题目限制了 1 <= num <= 3999，所以大可以花比较少的精力就把1 ~ 3999 运算所需的资源全部列出来然后逐一映射。
// 有在确定参数规模不大的情况下，一一列出所有可能情况也许反而是比较聪明、快速的做法。
// 时间复杂度：O(1)
// 空间复杂度：O(1)
func Q12Solution2(num int) string {
	m := map[int]string{0:"",1:"M",2:"MM",3:"MMM",}
	c := map[int]string{0:"",1:"C",2:"CC",3:"CCC",4:"CD",5:"D",6:"DC",7:"DCC",8:"DCCC",9:"CM"}
	x := map[int]string{0:"",1:"X",2:"XX",3:"XXX",4:"XL",5:"L",6:"LX",7:"LXX",8:"LXXX",9:"XC"}
	i := map[int]string{0:"",1:"I",2:"II",3:"III",4:"IV",5:"V",6:"VI",7:"VII",8:"VIII",9:"IX"}

	return m[num/1000] + c[(num%1000)/100] + x[(num%100)/10] + i[num%10]
}