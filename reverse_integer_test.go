package leetcode

import (
	"testing"
)

type ReverseIntegerCase struct {
	name   string
	num    int
	result int
}

func (c *ReverseIntegerCase) ok(r int) bool {
	return c.result == r
}

func createReverseIntegerTestCase(t *testing.T, c *ReverseIntegerCase) {
	t.Helper()
	result := reverse(c.num)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestReverse(t *testing.T) {
	cases := []*ReverseIntegerCase{
		{
			"case 1",
			123,
			321,
		},
		{
			"case 2",
			-123,
			-321,
		},
		{
			"case 3",
			120,
			21,
		},
		{
			"case 4",
			1534236469,
			0,
		},
	}

	for _, c := range cases {
		createReverseIntegerTestCase(t, c)
	}
}
