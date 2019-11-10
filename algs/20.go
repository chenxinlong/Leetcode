package algs

// 题20：https://leetcode-cn.com/problems/valid-parentheses/solution/valid-parentheses-fu-zhu-zhan-fa-by-jin407891080/

// 这道题太经典了，以致于一看到这道题首先想到的就是用 stack 或类 stack 来解。如果语言没有原生 stack 实现，其他能进行 fifo(first in, first out) 的数据结构也行。
// 时间复杂度：O(n)
// 空间复杂度：O(n)
// 知道用 stack，也知道大概思路。 但是没写过，所以漏掉各种 corner(edge) case，以致于提交了 5 次才通过所有 test case。
func Q20Solution1(s string) bool {
	// 如果 s 长度为奇数，则直接返回 false
	if len(s) % 2 != 0 {
		return false
	}

	stack := []string{}

	m := map[string]string{"]":"[", ")":"(", "}":"{"}

	for _, el := range s {
		elStr := string(el)
		// 如果当前元素是左括号，那么入栈
		// 如果当前元素是右括号，那么这个右括号应该和栈顶元素应该可以配对。如果配对则出栈，如果不配对则直接返回结果 false
		if elStr == "(" || elStr == "[" || elStr == "{" {
			stack = append(stack, elStr)
		} else {
			// 如果 s 一开始就是右括号，那么 stack 长度为 0，这里就会数组越界
			//if m[elStr] != stack[len(stack)-1] {
			//	return false
			//}

			if len(stack) == 0 {
				return false
			}
			if m[elStr] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	// 如果最后栈不为空，则说明有残余左括号未被匹配
	if len(stack) == 0 {
		return true
	}

	return false
}