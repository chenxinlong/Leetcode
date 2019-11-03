package algs

// 题6：https://leetcode.com/problems/zigzag-conversion/

// 假设 s = PAHNAPLSIIGYIR, numRow = 4, 想象出一个平面直角坐标系(y轴向下)
//   0 1 2 3 4 5 6 7 8
//  .---------------> x
// 0|P     I     N
// 1|A   L S   I G
// 2|Y A   H R
// 3|P     I
// 4|
//  ↓
//  y
// 我们把坐标系上的每一个坐标的值写入到二维数组，然后遍历二维数组的每一行每一列，把所有不为0的值拼接起来。二维数组如下：
// m = [
//   [P A Y P],
//   [0 0 A 0],
//   [0 L 0 0],
//   [I S H I],
//   [0 0 R 0],
//   [0 I 0 0],
//   [N G 0 0]
// ]
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func Q6Solution1(s string, numRows int) string {
	l := len(s)
	if l < numRows {
		return s
	}

	result := ""
	m := [][]uint8{}
	k := 0
	for i:=0; i<l && k<l; i++ {
		col := []uint8{}
		for j:=0; j<numRows; j++ {
			if k < l && (i%(numRows-1) == 0 || (i+j)%(numRows-1) == 0) {
				col = append(col, s[k])
				k++
			} else {
				col = append(col, 0)
			}
		}
		m = append(m, col)
	}

	for row:=0;row<numRows;row++ {
		for col :=0;col<len(m);col++ {
			if m[col][row] != 0 {
				result += string(m[col][row])
			}
		}
	}

	return result
}

// 解法1的优化版：
// 假设 s = PAHNAPLSIIGYIR, numRow = 4, 和解法1一样，我们同样是想象出一个平面直角坐标系(y轴向下).
// 声明一个指向当前行的指针 currRow (并非真正意义上的指针，而是存储了当前所处行号的变量)
//       0 1 2 3 4 5 6 7 8
//      .---------------> x
//     0|P     I     N
//   → 1|A   L S   I G
//     2|Y A   H R
//     3|P     I
//     4|
//      ↓
//      y
// 遍历字符串每个字符，每次迭代都把当前字符拼接到 string[currRow] 后面, 然后 currRow 随即移动到下(或上)一行。当 currRow == 0，或者 currRow == 3 改变方向。
// 解法1是把坐标系中的每一坐标都走一遍，并全部计入二维数组之后开始拼接结果。该解法是移动 curRow (+1 或 -1) ，这个思路挺巧妙的,可以记住。
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func Q6Solution2(s string, numRows int) string {
	l := len(s)
	if numRows < 2 || l <= numRows {
		return s
	}

	rows := map[int]string{}
	currRow := 0
	goingDown := false
	for _, char := range s {
		rows[currRow] += string(char)

		// 如果当前行是第0行或最后一行(即应该拐角的那行)，那么改变方向。
		// 但是要注意 numRows == 1 这种情况，每一次该判断都会为 true。这种情况在函数开头就直接过滤掉。 ---反思：每次写 if 的时候都要考虑边缘用例、极端情况。
		if currRow == 0 || currRow == numRows-1 {
			goingDown = !goingDown
		}

		// goingDown == true, currRow 向下移动
		// goingDown == false, currRow 向上移动
		if goingDown {
			currRow += 1
		} else {
			currRow -= 1
		}
	}

	result := ""
	// map 的遍历是无序随机的，所以这里不能用 range 遍历 rows
	//for _, row := range rows {
	//	result += row
	//}
	for i:=0; i < numRows; i++ {
		result+=rows[i]
	}

	return result
}