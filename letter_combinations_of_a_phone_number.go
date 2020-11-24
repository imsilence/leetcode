package leetcode

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。



示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*/

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	numMap := map[rune]string{
		'0': "_",
		'1': "!@#",
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
		'*': "+",
		'#': "",
	}

	texts := make([]string, len(digits))
	ceils := make([]int, len(digits))
	for i, num := range digits {
		texts[i] = numMap[num]
		ceils[i] = len(numMap[num])
	}

	indexs := make([]int, len(digits))
	next := func() bool {
		indexs[0]++
		flag := 0
		for i, ceil := range ceils {
			indexs[i] += flag
			if indexs[i] >= ceil {
				indexs[i] = 0
				flag = 1
			} else {
				flag = 0
			}

		}
		return flag == 0
	}
	text := func() string {
		result := make([]byte, len(indexs))
		for i, index := range indexs {
			result[i] = texts[i][index]
		}
		return string(result)
	}

	result := make([]string, 0)
	for {
		result = append(result, text())
		if !next() {
			break
		}
	}
	return result
}
