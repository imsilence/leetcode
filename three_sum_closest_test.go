package leetcode

import (
	"testing"
)

type ThreeSumClosestCase struct {
	name   string
	list   []int
	target int
	result int
}

func (c *ThreeSumClosestCase) ok(r int) bool {
	return c.result == r
}

func createThreeSumClosestTestCase(t *testing.T, c *ThreeSumClosestCase) {
	t.Helper()
	result := threeSumClosest(c.list, c.target)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestThreeSumClosest(t *testing.T) {
	cases := []*ThreeSumClosestCase{
		{
			"case 1",
			[]int{-1, 2, 1, -4},
			1,
			2,
		},
	}

	for _, c := range cases {
		createThreeSumClosestTestCase(t, c)
	}
}
