package leetcode

import (
	"testing"
)

type RegularExpressionMatchingCase struct {
	name   string
	text   string
	expr   string
	result bool
}

func (c *RegularExpressionMatchingCase) ok(r bool) bool {
	return c.result == r
}

func createRegularExpressionMatchingCase(t *testing.T, c *RegularExpressionMatchingCase) {
	t.Helper()
	result := isMatch(c.text, c.expr)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestIsMatch(t *testing.T) {
	cases := []*RegularExpressionMatchingCase{
		{
			"case 1",
			"aa",
			"a",
			false,
		},
		{
			"case 2",
			"aa",
			"a*",
			true,
		},
		{
			"case 3",
			"ab",
			"a.",
			true,
		},
		{
			"case 4",
			"aab",
			"c*a*b",
			true,
		},
		{
			"case 5",
			"mississippi",
			"mis*is*p*.",
			false,
		},
		{
			"case 6",
			"ab",
			".*c",
			false,
		},
	}

	for _, c := range cases {
		createRegularExpressionMatchingCase(t, c)
	}
}
