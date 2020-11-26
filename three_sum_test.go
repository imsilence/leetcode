package leetcode

import (
	"testing"
)

type ThreeSumCase struct {
	name   string
	list   []int
	result [][]int
}

func (c *ThreeSumCase) ok(r [][]int) bool {
	return Int2Equals(c.result, r, true)
}

func createThreeSumTestCase(t *testing.T, c *ThreeSumCase) {
	t.Helper()
	result := threeSum(c.list)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestThreeSum(t *testing.T) {
	cases := []*ThreeSumCase{
		{
			"case 1",
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				{-1, -1, 2},
				{1, 0, -1},
			},
		},
	}

	for _, c := range cases {
		createThreeSumTestCase(t, c)
	}
}
