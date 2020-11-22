package leetcode

import "testing"

type LongestPalindromicSubstringCase struct {
	name   string
	text   string
	result []string
}

func (c *LongestPalindromicSubstringCase) ok(r string) bool {
	for _, rs := range c.result {
		if rs == r {
			return true
		}
	}

	return false
}

func createLongestPalindromicSubstringTestCase(t *testing.T, c *LongestPalindromicSubstringCase) {
	t.Helper()
	result := longestPalindrome(c.text)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}

}

func TestLongestPalindrome(t *testing.T) {
	cases := []*LongestPalindromicSubstringCase{
		{
			"case 1",
			"babad",
			[]string{"bab", "aba"},
		},
		{
			"case 2",
			"cbbd",
			[]string{"bb"},
		},
	}
	for _, c := range cases {
		createLongestPalindromicSubstringTestCase(t, c)
	}
}
