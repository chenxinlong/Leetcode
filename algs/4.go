package algs

import (
	"math"
	"sort"
)

// 题4：https://leetcode.com/problems/median-of-two-sorted-arrays/
// todo : 这题至少还有 2 种解法，后面补上

// 首先想到的解法：合并数组重排，然后取中位数。但是不满足题目要求的 O(log(m+n)) 时间复杂度
// 时间复杂度：排序算法的时间复杂度
// 空间复杂度：排序算法的空间复杂度
func Q4Solution1(nums1 []int, nums2 []int) float64 {
	numsTotal := append(nums1, nums2...)
	sort.Ints(numsTotal)

	l := len(numsTotal)
	if l%2==1 {
		return float64(numsTotal[l/2])
	}

	result := float64(numsTotal[l/2]+numsTotal[l/2-1])/2
	return result
}

// 该解法借鉴了别人的题解，但是这道题leetcode给的题解不是那么好理解
// 比较通俗的理解可以看 https://medium.com/@hazemu/finding-the-median-of-2-sorted-arrays-in-logarithmic-time-1d3f2ecbeb46
func Q4Solution2(nums1 []int, nums2 []int) float64 {
	// 我们重新定义 a, b 数组, 其中 a 表示元素比较少的数组， b 表示元素比较多的数组
	a, b, aLen, bLen := nums1, nums2, len(nums1), len(nums2)
	if aLen > bLen {
		a, b, aLen, bLen = b, a, bLen, aLen
	}

	// 1.假设 aLen + bLen 为奇数，则a,b合并后的数组的的左半部分长度为 (aLen + bLen)/2并向上取整
	// 2.假设 aLen + bLen 为偶数，则a,b合并后的数组的的左半部分长度为 (aLen + bLen)/2
	// 在代码中，我们直接 (aLen + bLen + 1) / 2，就可以避免多写一些奇偶数判断的逻辑。
	leftHalfLen := (aLen + bLen + 1) / 2

	// 比较短的数组，最少为 leftHalf 贡献 0 个元素，最多为 leftHalf 贡献全部元素
	aMinContribute := 0
	aMaxContribute := aLen

	// 题目要求的复杂度 O(log(m+n)) ，实际上是这里的二分查找。二分查找是一个如果不仔细判断，很容易就不小心出现数组越界的算法。
	// 假定 a 贡献 m 个元素给 leftHalf，那么 b 则贡献 leftHalfLen - m 个元素
	for aMinContribute <= aMaxContribute {
		// aContribute 原本应该是 (aMaxContribute + aMinContribute)/2，但是用这个公式在某些 edge case 下会导致后面流程中出现数组越界
		aContribute := aMinContribute + ((aMaxContribute - aMinContribute) / 2)
		bContribute := leftHalfLen - aContribute

		if aContribute > 0 && a[aContribute-1] > b[bContribute]  {
			// a 贡献元素没有那么多，要减少一个
			aMaxContribute = aContribute - 1
		} else if aContribute < aLen  && b[bContribute-1] > a[aContribute] {
			// a 贡献元素不止这么少，要增加一个
			aMinContribute = aContribute + 1
		} else {
			var leftHalfEnd float64
			// 如果 a 贡献 0 个元素，那么 leftHalf 最后一个元素就是 b 贡献的最后一个元素
			// 如果 b 贡献 0 个元素，那么 leftHalf 最后一个元素就是 a 贡献的最后一个元素
			// 如果 a,b 各自都贡献了元素，那么 leftHalf 最后一个元素就是 a,b贡献的最后一个元素的较大者。
			if aContribute == 0 {
				leftHalfEnd = float64(b[bContribute - 1])
			} else {
				if bContribute == 0 {
					leftHalfEnd = float64(a[aContribute-1])
				} else {
					leftHalfEnd = math.Max(float64(a[aContribute-1]), float64(b[bContribute-1]))
				}
			}

			// 如果 nums1 ∪ nums2 的数组长度是奇数，则中位数就是 leftHalf 的最后一个元素
			if (aLen + bLen) & 1 == 1 {
				return leftHalfEnd
			}

			// 如果 nums1 ∪ nums2 的数组长度是偶数，那么我们还得算出 rightHalf 的第一个元素，与 leftHalf 最后一个元素做均值。
			var rightHalfStart float64

			// 如果 a 贡献了全部元素给 leftHalf，那么 rightHalf 第一个元素就是 b 贡献的最后一个元素的下一个元素
			// 如果 b 贡献了全部元素给 leftHalf，那么 rightHalf 第一个元素就是 a 贡献的最后一个元素的下一个元素
			// 如果 a,b 各自都只贡献了自己的一部分元素给 leftHalf, 那么 a,b 各自最后贡献的元素的后一个元素的较大者就是 rightHalf 的第一个元素
			if aContribute == aLen {
				rightHalfStart = float64(b[bContribute])
			} else {
				if bContribute == bLen {
					rightHalfStart = float64(a[aContribute])
				} else {
					rightHalfStart = math.Max(float64(b[bContribute]), float64(a[aContribute]))
				}
			}

			return (leftHalfEnd + rightHalfStart)/2
		}
	}

	return -1
}

