package leetcode

/*
5. 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/
func longestPalindrome(s string) string {
	result := ""

	palindrome := func(start, end int) string {
		for start >= 0 && end < len(s) {
			if s[start] != s[end] {
				break
			}
			start--
			end++

		}
		return s[start+1 : end]
	}

	for i := 0; i < len(s); i++ {
		if txt := palindrome(i, i); len(txt) > len(result) {
			result = txt
		}
		if txt := palindrome(i, i+1); len(txt) > len(result) {
			result = txt
		}
	}

	return result
}
