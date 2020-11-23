package leetcode

/*
14. 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	char := func(text string, index int) byte {
		if len(text) <= index {
			return 0
		}
		return text[index]
	}

	match := func(index int) bool {
		ch := char(strs[0], index)

		for _, text := range strs {
			if c := char(text, index); c == 0 || ch != c {
				return false
			}
		}

		return true
	}

	index := 0
	for match(index) {
		index++
	}
	return strs[0][:index]
}
