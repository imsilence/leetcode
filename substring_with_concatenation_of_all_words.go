package leetcode

/*
30. 串联所有单词的子串
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。



示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]
*/

func findSubstring(s string, words []string) []int {
	indexMap := map[int]string{}
	indexs := make([]int, 0, len(words))
	wordsMap := make(map[string]int)

	for _, word := range words {
		wordsMap[word]++
	}

NEXT:
	for i := 0; i < len(s); i++ {
	WORD:
		for word := range wordsMap {
			if word[0] != s[i] || len(s)-i < len(word) {
				continue
			}
			for j := 1; j < len(word); j++ {
				if word[j] != s[j+i] {
					continue WORD
				}
			}
			indexMap[i] = word
			indexs = append(indexs, i)
			continue NEXT
		}
	}

	if len(indexs) < len(words) {
		return nil
	}

	sortInts := func(list []int) {
		for i := 0; i < len(list)-1; i++ {
			for j := 0; j < len(list)-i-1; j++ {
				if list[j] > list[j+1] {
					list[j], list[j+1] = list[j+1], list[j]
				}
			}
		}
	}

	equal := func(words map[string]int) bool {
		if len(wordsMap) != len(words) {
			return false
		}
		for k, c := range wordsMap {
			if cc, ok := words[k]; !ok || c != cc {
				return false
			}
		}
		return true
	}

	sortInts(indexs)
	result := []int{}
INDEX:
	for _, v := range indexs {
		findWords := map[string]int{}
		for i := 0; i < len(words); i++ {
			word, ok := indexMap[v+i*len(words[0])]
			if !ok {
				continue INDEX
			}
			findWords[word]++
		}
		if equal(findWords) {
			result = append(result, v)
		}
	}

	return result
}
