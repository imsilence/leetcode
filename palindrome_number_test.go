package leetcode

import (
	"testing"
)

type PalindromeNumberCase struct {
	name   string
	num    int
	result bool
}

func (c *PalindromeNumberCase) ok(r bool) bool {
	return c.result == r
}

func createPalindromeNumberCaseTestCase(t *testing.T, c *PalindromeNumberCase) {
	t.Helper()
	result := isPalindrome(c.num)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := []*PalindromeNumberCase{
		{
			"case 1",
			121,
			true,
		},
		{
			"case 2",
			-121,
			false,
		},
		{
			"case 3",
			10,
			false,
		},
	}

	for _, c := range cases {
		createPalindromeNumberCaseTestCase(t, c)
	}
}
