package leetcode

import "fmt"

/*
18. 四数之和
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/
func fourSum(nums []int, target int) [][]int {
	exists := make(map[string]bool)
	result := [][]int{}
	sort := func(nums []int) {
		for i := 0; i < len(nums)-1; i++ {
			for j := 0; j < len(nums)-1-i; j++ {
				if nums[j] > nums[j+1] {
					nums[j], nums[j+1] = nums[j+1], nums[j]
				}
			}
		}
	}
	key := func(nums ...int) string {
		sort(nums)
		return fmt.Sprintf("%d_%d_%d_%d", nums[0], nums[1], nums[2], nums[3])
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for m := j + 1; m < len(nums); m++ {
				for n := m + 1; n < len(nums); n++ {
					if nums[i]+nums[j]+nums[m]+nums[n] == target {
						k := key(nums[i], nums[j], nums[m], nums[n])
						if _, ok := exists[k]; !ok {
							exists[k] = true
							result = append(result, []int{nums[i], nums[j], nums[m], nums[n]})
						}
					}
				}
			}
		}
	}
	return result
}
