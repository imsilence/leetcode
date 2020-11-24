package leetcode

import "testing"

type ImplementStrStrCase struct {
	name    string
	text    string
	subtext string
	result  int
}

func (c *ImplementStrStrCase) ok(r int) bool {
	return c.result == r
}

func createImplementStrStrTestCase(t *testing.T, c *ImplementStrStrCase) {
	t.Helper()
	result := strStr(c.text, c.subtext)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestStrStr(t *testing.T) {
	cases := []*ImplementStrStrCase{
		{
			"case 1",
			"hello",
			"ll",
			2,
		},
		{
			"case 2",
			"aaaaa",
			"bba",
			-1,
		},
	}

	for _, c := range cases {
		createImplementStrStrTestCase(t, c)
	}

}
