package leetcode

import "testing"

type LongestValidPaenthesesCase struct {
	name   string
	text   string
	result int
}

func (c *LongestValidPaenthesesCase) ok(r int) bool {
	return c.result == r
}

func createLongestValidPaenthesesTestCase(t *testing.T, c *LongestValidPaenthesesCase) {
	t.Helper()
	result := longestValidParentheses(c.text)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestLongestValidPaentheses(t *testing.T) {
	cases := []*LongestValidPaenthesesCase{
		{
			"case 1",
			"(()",
			2,
		},
		{
			"case 2",
			")()())",
			4,
		},
		{
			"case 3",
			"()(()",
			2,
		},
		{
			"case 4",
			"()(()(((",
			2,
		},
		{
			"case 5",
			"((()()(()((()",

			4,
		},
		{
			"case 5",
			")(()((((())",
			4,
		},
	}

	for _, c := range cases {
		createLongestValidPaenthesesTestCase(t, c)
	}

}
