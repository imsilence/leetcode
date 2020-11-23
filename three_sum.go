package leetcode

/*
15. 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
func threeSum(nums []int) [][]int {
	lines := map[[3]int]bool{}
	numMap := map[int]int{}

	key := func(a, b, c int) [3]int {
		if b > a {
			a, b = b, a
		}
		if c > a {
			a, c = c, a
		}

		if c > b {
			b, c = c, b
		}
		return [3]int{a, b, c}
	}

	for _, num := range nums {
		numMap[num]++
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			r := nums[i] + nums[j]
			if v, ok := numMap[-r]; !ok || ((-r == nums[i] || -r == nums[j]) && v < 2) || (-r == nums[i] && -r == nums[j] && v < 3) {
				continue
			}
			k := key(nums[i], nums[j], -r)
			if _, ok := lines[k]; ok {
				continue
			}

			lines[k] = true
		}
	}

	result := make([][]int, 0, len(lines))
	for line := range lines {
		t := line
		result = append(result, t[:])
	}
	return result
}
