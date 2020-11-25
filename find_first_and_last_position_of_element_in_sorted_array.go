package leetcode

/*
34. 在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/
func searchRange(nums []int, target int) []int {
	binarySearch := func() int {
		start, end := 0, len(nums)-1
		for start <= end {
			middle := (start + end) / 2
			if nums[middle] == target {
				return middle
			} else if nums[middle] < target {
				start = middle + 1
			} else {
				end = middle - 1
			}
		}
		return -1
	}

	index := binarySearch()
	if index == -1 {
		return []int{-1, -1}
	}
	start, end := index, index
	for start >= 0 && nums[start] == target {
		start--
	}
	for end < len(nums) && nums[end] == target {
		end++
	}

	return []int{start + 1, end - 1}
}
