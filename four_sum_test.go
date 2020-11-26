package leetcode

import (
	"testing"
)

type FourSumCase struct {
	name   string
	list   []int
	target int
	result [][]int
}

func (c *FourSumCase) ok(r [][]int) bool {
	return Int2Equals(c.result, r, true)
}

func createFourSumTestCase(t *testing.T, c *FourSumCase) {
	t.Helper()
	result := fourSum(c.list, c.target)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestFourSum(t *testing.T) {
	cases := []*FourSumCase{
		{
			"case 1",
			[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
	}

	for _, c := range cases {
		createFourSumTestCase(t, c)
	}
}
