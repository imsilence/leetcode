package leetcode

/*
35. 搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0
*/
func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	} else if target == nums[len(nums)-1] {
		return len(nums) - 1
	} else if target > nums[len(nums)-1] {
		return len(nums)
	} else {
		binarySearch := func() int {
			start, end := 0, len(nums)
			for start <= end {
				m := (start + end) / 2
				if nums[m] == target {
					return m
				} else if nums[m] > target {
					end = m - 1
				} else {
					start = m + 1
				}
			}
			return end + 1
		}
		return binarySearch()

	}
}
