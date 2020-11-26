package leetcode

/*
40. 组合总和 II
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]40. 组合总和 II
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
*/

func combination2(start int, candidates []int, target int, line []int, lines [][]int) [][]int {
	i := start
	for i < len(candidates) {
		v := candidates[i]
		nline := append([]int{}, line...)
		ntarget := target - v
		if target < 0 {
			break
		} else if ntarget == 0 {
			nline = append(nline, v)
			lines = append(lines, nline)
		} else {
			nline = append(nline, v)
			lines = combination2(i+1, candidates, ntarget, nline, lines)
		}
		for i < len(candidates) && candidates[i] == v {
			i++
		}
	}
	return lines
}

func combinationSum2(candidates []int, target int) [][]int {
	lines := make([][]int, 0)

	ints := func(list []int) {
		for i := 0; i < len(list)-1; i++ {
			for j := 0; j < len(list)-i-1; j++ {
				if list[j] > list[j+1] {
					list[j], list[j+1] = list[j+1], list[j]
				}
			}
		}
	}

	ints(candidates)

	lines = combination2(0, candidates, target, []int{}, lines)

	int2str := func(k int) string {
		m := map[int]string{
			1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 0: "0",
		}
		if k == 0 {
			return "0"
		}
		s := ""
		for k != 0 {
			s += m[k%10]
			k /= 10
		}
		return s
	}

	ints2str := func(list []int) string {
		s := "-"
		for _, v := range list {
			s += int2str(v)
			s += "-"
		}
		return s
	}

	key := func(list []int) string {
		ints(list)
		return ints2str(list)
	}

	result := make([][]int, 0)
	counts := make(map[string]bool)

	for i, line := range lines {
		k := key(line)
		if _, ok := counts[k]; !ok {
			counts[k] = true
			result = append(result, lines[i])
		}
	}
	return result
}
