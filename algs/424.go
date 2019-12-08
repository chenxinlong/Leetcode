package algs

// 题 424 : https://leetcode.com/problems/longest-repeating-character-replacement/

// ABBAAAABB, 把这个字符串变成全部字符一样的字符串，我们选择把非 A 变为 A 是最省力的，只要变动 4 个字符。
// 换成子串再来一遍：
// A  : 无需变
// AB : 变谁都行, 变1次
// ABB : 把 A 变成 B,变1次
// ABBA : 变谁都行,变2次
// ABBAA : 把 B 变成 A ,变2次
// ABBAAA: 把 B 变成 A , 变2次
// ABBAAAA: 把 B 变成 A , 变2次
// ABBAAAAB: 把 B 变成 A ,变3次
// ABBAAAABB: 把 B 变成 A , 变4次

func Q424Solution1(s string, k int) int {
	return 0
}