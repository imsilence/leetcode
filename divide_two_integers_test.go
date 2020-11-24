package leetcode

import "testing"

type DivideTwoIntegersCase struct {
	name   string
	left   int
	right  int
	result int
}

func (c *DivideTwoIntegersCase) ok(r int) bool {
	return c.result == r
}

func createDivideTwoIntegersTestCase(t *testing.T, c *DivideTwoIntegersCase) {
	t.Helper()
	result := divide(c.left, c.right)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestDivide(t *testing.T) {
	cases := []*DivideTwoIntegersCase{
		{
			"case 1",
			10,
			3,
			3,
		},
		{
			"case 2",
			7,
			-3,
			-2,
		},

		{
			"case 3",
			-2147483648,
			-1,
			2147483647,
		},
	}

	for _, c := range cases {
		createDivideTwoIntegersTestCase(t, c)
	}

}
