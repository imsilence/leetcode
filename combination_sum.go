package leetcode

/*
39. 组合总和
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]
示例 2：

输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]


提示：

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都是独一无二的。
1 <= target <= 500
*/

func combination(candidates []int, target int, line []int, lines [][]int) [][]int {
	for _, v := range candidates {
		nline := append([]int{}, line...)
		ntarget := target - v
		if target < 0 {
			break
		} else if target == 0 {
			nline = append(nline, v)
			lines = append(lines, line)
			break
		} else {
			nline = append(nline, v)
			lines = combination(candidates, ntarget, nline, lines)
		}
	}
	return lines
}

func combinationSum(candidates []int, target int) [][]int {
	lines := make([][]int, 0)
	lines = combination(candidates, target, []int{}, lines)

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

	ints2str := func(line []int) string {
		s := "-"
		for _, v := range line {
			s += int2str(v)
			s += "-"
		}
		return s
	}

	key := func(line []int) string {
		for i := 0; i < len(line)-1; i++ {
			for j := 0; j < len(line)-i-1; j++ {
				if line[j] > line[j+1] {
					line[j], line[j+1] = line[j+1], line[j]
				}
			}
		}
		return ints2str(line)
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
