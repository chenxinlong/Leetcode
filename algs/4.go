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
// 时间复杂度：O(log(min(m,n)))
// 空间复杂度：O(1)
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

	// 比较短的数组最少为 leftHalf 贡献 0 个元素，最多为 leftHalf 贡献全部元素
	aMinContribute := 0
	aMaxContribute := aLen

	// 假设结果是 a,b 各自贡献 x,y 个元素给 leftHalf，那么这个算法本质就是在根据以下几个性质来不断验证和调整 x 和 y :
	// 1. x + y = len(a ∪ b)/2
	// 2. median = max(a[x-1], b[y-1])
	// 3. a[x] > median, b[y] > median
	for aMinContribute <= aMaxContribute {
		// 假定 a 贡献 m 个元素给 leftHalf，那么 b 则贡献 leftHalfLen - m 个元素
		// aContribute 原本应该是 (aMaxContribute + aMinContribute)/2，但是用这个公式在某些 edge case 下会导致后面流程中出现数组越界
		aContribute := aMinContribute + ((aMaxContribute - aMinContribute) / 2)
		bContribute := leftHalfLen - aContribute

		// 如果 a[x-1] > b[y-1]，那么说明 a[x-1] 是中位数(leftHalf最后一个元素/leftHalf最大元素)。
		// 这样一来， rightHalf 不可能有比 a[x-1] 更小的元素，也就是处于 rightHalf 的 b[y] 不可能比 a[x-1] 小。如果有则说明 a[x-1] 一定在 rightHalf, 也就是 a 不可能贡献 x 这么多个元素给 leftHalf，那么我们就得把 a[x-1] 元素剔除掉再次试错
		if aContribute > 0 && a[aContribute-1] > b[bContribute]  {
			aMaxContribute = aContribute - 1

		// 如果 b[y-1] > a[x-1]，那么说明 b[y-1] 是中位数(leftHalf最后一个元素/leftHalf最大元素)。
		// 这样一来，rightHalf 不可能有笔 b[y-1] 更小的元素，也就是处于 rightHalf 的 a[x] 不可能比 b[y-1] 小。如果有则说明 a[x] 一定在 leftHalf，也就是 a 不可能贡献 x 这么少个元素给 leftHlaf, 那么我们就得把 a[x] 元素假如到 leftHalf 再次试错
		} else if aContribute < aLen  && b[bContribute-1] > a[aContribute] {
			aMinContribute = aContribute + 1

		// 如果进入最后一个 else，表明我们已经得到正确的 x 和 y, 那么接下来就是求中位数
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

