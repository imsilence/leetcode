package leetcode

import (
	"testing"
)

type LongestSubstringWithoutRepeatingCharactersTestCase struct {
	name   string
	text   string
	result int
}

func (c *LongestSubstringWithoutRepeatingCharactersTestCase) ok(r int) bool {
	return c.result == r
}

func createLongestSubstringWithoutRepeatingCharactersTestCase(t *testing.T, c *LongestSubstringWithoutRepeatingCharactersTestCase) {
	t.Helper()
	result := lengthOfLongestSubstring(c.text)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []*LongestSubstringWithoutRepeatingCharactersTestCase{
		{
			"case 1",
			"abcabcbb",
			3,
		},
		{
			"case 2",
			"bbbbb",
			1,
		},
		{
			"case 3",
			"pwwkewx",
			4,
		},
	}
	for _, c := range cases {
		createLongestSubstringWithoutRepeatingCharactersTestCase(t, c)
	}
}
