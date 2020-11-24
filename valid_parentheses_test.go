package leetcode

import (
	"testing"
)

type ValidParenthesesCase struct {
	name   string
	text   string
	result bool
}

func (c *ValidParenthesesCase) ok(r bool) bool {
	return c.result == r
}

func createValidParenthesesTestCase(t *testing.T, c *ValidParenthesesCase) {
	t.Helper()
	result := isValid(c.text)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestIsValid(t *testing.T) {
	cases := []*ValidParenthesesCase{
		{
			"case 1",
			"()",
			true,
		},
		{
			"case 2",
			"()[]{}",
			true,
		},
		{
			"case 3",
			"(]",
			false,
		},
		{
			"case 4",
			"([)]",
			false,
		},
		{
			"case 5",
			"{[]}",
			true,
		},
	}

	for _, c := range cases {
		createValidParenthesesTestCase(t, c)
	}
}
