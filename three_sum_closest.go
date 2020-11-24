package leetcode

/*
16. 最接近的三数之和
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。



示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。


提示：

3 <= nums.length <= 10^3
-10^3 <= nums[i] <= 10^3
-10^4 <= target <= 10^4
*/
func threeSumClosest(nums []int, target int) int {
	result := 0
	minDiff := 1<<31 - 1
	for x := 0; x < len(nums); x++ {
		for y := x + 1; y < len(nums); y++ {
			for z := y + 1; z < len(nums); z++ {
				sum := nums[x] + nums[y] + nums[z]
				diff := sum - target
				if diff == 0 {
					return sum
				}
				if diff < 0 {
					diff = -diff
				}
				if diff < minDiff {
					minDiff = diff
					result = sum
				}
			}
		}
	}
	return result
}
