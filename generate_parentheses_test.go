package leetcode

import (
	"testing"
)

type GenerateParenthesesCase struct {
	name   string
	num    int
	result []string
}

func (c *GenerateParenthesesCase) ok(r []string) bool {
	return StringEquals(c.result, r, true)
}

func createGenerateParenthesesTestCase(t *testing.T, c *GenerateParenthesesCase) {
	t.Helper()
	result := generateParenthesis(c.num)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestGenerateParenthesis(t *testing.T) {
	cases := []*GenerateParenthesesCase{
		{
			"case 1",
			1,
			[]string{
				"()",
			},
		},
		{
			"case 2",
			3,
			[]string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
	}

	for _, c := range cases {
		createGenerateParenthesesTestCase(t, c)
	}

}
