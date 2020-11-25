package leetcode

/*
31. 下一个排列
实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须 原地 修改，只允许使用额外常数空间。



示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]
示例 2：

输入：nums = [3,2,1]
输出：[1,2,3]
示例 3：

输入：nums = [1,1,5]
输出：[1,5,1]
示例 4：

输入：nums = [1]
输出：[1]


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 100
*/
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	isDesc := func() (bool, int) {
		for i := len(nums) - 1; i > 0; i-- {
			if nums[i-1] < nums[i] {
				return false, i - 1
			}
		}
		return true, 0
	}

	if desc, index := isDesc(); desc {
		for i := 0; i < len(nums)/2; i++ {
			nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
		}
	} else {
		swapIndex := index + 1
		for i := index + 2; i < len(nums); i++ {
			if nums[i] > nums[index] && nums[i] < nums[swapIndex] {
				swapIndex = i
			}
		}
		nums[index], nums[swapIndex] = nums[swapIndex], nums[index]
		for j := 0; j < len(nums)-1-index-1; j++ {
			for i := index + 1; i < len(nums)-1-j; i++ {
				if nums[i] > nums[i+1] {
					nums[i], nums[i+1] = nums[i+1], nums[i]
				}
			}
		}
	}
}
