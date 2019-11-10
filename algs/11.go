package algs

// 题11：https://leetcode.com/problems/container-with-most-water/
// 这题转化一下，其实就是求 height 数组里，哪2个元素相乘最大

// Brute force : 无脑暴力解法，就是把所有可能的 长 * 宽 都计算一遍，取最大值
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func Q11Solution1(height []int) int {
	l := len(height)
	maxArea := 0
	for i:=0; i < l; i++ {
		for j:=i+1; j<l; j++ {
			// 仅考虑矮边
			iSqr := (j-i) * height[i]
			jSqr := (j-i) * height[j]
			if (iSqr >= jSqr) && jSqr >= maxArea {
				maxArea = jSqr
			} else if (jSqr >= iSqr) && iSqr >= maxArea {
				maxArea = iSqr
			}
		}
	}

	return maxArea
}

// 双指针 : 把双指针解法写出来并不难。难的是，双指针为什么能解这题？看了下讨论区，大部分人的疑问和我一样，就是：这题凭什么能用双指针解!!!
// 用暴力法解，能保证覆盖到每一个可能出现的结果，然后在所有可能结果中取 maxArea。
// 用双指针解，一次头(或尾)指针的移动，就会有一批可能性被排除掉。那些被排除掉(即没有被覆盖到) 的结果，怎么证明它们当中不会存在 maxArea ？ 或者说凭什么认为把它们排除掉是安全的？
// 解：
//  1. 初始状态 {left:0, right:length-1} 包含了所有可能性结果，所以是安全的
//  2. 假设指向较长的元素的指针是 "长指针"，否则是 "短指针"， 按照题解，我们每次都移动短直针，那么怎么判定每次移动短直针后排除掉的结果都是安全的？
//  todo
//
// 双指针问题：每次迭代，1.如何判定移动 left 还是 right ?  2. 每一次迭代 left, right 都是向中间移动吗，有没有可能出现在某次迭代过程中其中一个指针向外侧移动的情况？
// 时间复杂度：O(n)
// 空间复杂度：O(1)
// 已知解法，去证明结果是对的。而不是根据题目，想出如何得出这个解法。 这道题做得实在有点憋屈....
func Q11Solution2(height []int) int {
	left, right, maxArea := 0, len(height)-1, 0
	for left < right {
		currentArea := min(height[left], height[right]).(int) * (right-left)
		maxArea = max(currentArea, maxArea).(int)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}