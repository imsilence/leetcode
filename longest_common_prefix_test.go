package leetcode

import "testing"

type LongestCommonPrefixCase struct {
	name   string
	texts  []string
	result string
}

func (c *LongestCommonPrefixCase) ok(r string) bool {
	return c.result == r
}

func createLongestCommonPrefixTestCase(t *testing.T, c *LongestCommonPrefixCase) {
	t.Helper()
	result := longestCommonPrefix(c.texts)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	cases := []*LongestCommonPrefixCase{
		{
			"case 1",
			[]string{"flower", "flow", "flight"},
			"fl",
		},
		{
			"case 2",
			[]string{"dog", "racecar", "car"},
			"",
		},
		{
			"case 3",
			[]string{"", "", ""},
			"",
		},
	}

	for _, c := range cases {
		createLongestCommonPrefixTestCase(t, c)
	}

}
