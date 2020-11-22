package leetcode

/*
4. 寻找两个正序数组的中位数
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？



示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
示例 4：

输入：nums1 = [], nums2 = [1]
输出：1.00000
示例 5：

输入：nums1 = [2], nums2 = []
输出：2.00000

*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	start1, end1 := 0, len(nums1)-1
	start2, end2 := 0, len(nums2)-1
	min, max := 0, 0

	gte := func(index1, index2 int) (bool, bool, bool) {
		if index1 >= 0 && index1 < len(nums1) {
			if index2 >= 0 && index2 < len(nums2) {
				return nums1[index1] >= nums2[index2], false, false
			}
			return true, false, true
		}

		return false, true, false
	}
	for start1 <= end1 || start2 <= end2 {
		if ok, overflow1, overflow2 := gte(start1, start2); ok {
			if !overflow2 {
				min = nums2[start2]
				start2++
			} else {
				min = nums1[start1]
				start1++
			}
		} else {
			if !overflow1 {
				min = nums1[start1]
				start1++
			} else {
				min = nums2[start2]
				start2++
			}
		}

		if ok, overflow1, overflow2 := gte(end1, end2); ok {
			if !overflow1 {
				max = nums1[end1]
				end1--
			} else {
				max = nums2[end2]
				end2--
			}
		} else {
			if !overflow2 {
				max = nums2[end2]
				end2--
			} else {
				max = nums1[end1]
				end1--

			}
		}
	}
	return (float64(min) + float64(max)) / 2
}
