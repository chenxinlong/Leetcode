package algs


func DiStringMatch(S string) []int {
	res := make([]int, len(S))
	i:=0
	j:=len(S)
	for _, byte := range S {
		letter := string(byte)
		if letter == "I" {
			res = append(res, i)
			i++
		}
		if letter == "D" {
			res = append(res, j)
			j--
		}
	}
	return append(res, i)
}


/**
 * [analysis]
 * 归纳法，看例子归纳规律。 I 升序，D 降序，有个 leading 0 ，所以最后面要 append i (append j 也行，最后因为 i==j)
 */