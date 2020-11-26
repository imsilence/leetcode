package leetcode

import (
	"testing"
)

type TwoSumCase struct {
	name   string
	list   []int
	target int
	result []int
}

func (c *TwoSumCase) ok(r []int) bool {
	return IntEquals(c.result, r, true)
}

func createTwoSumTestCase(t *testing.T, c *TwoSumCase) {
	t.Helper()
	result := twoSum(c.list, c.target)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestTwoSum(t *testing.T) {
	cases := []*TwoSumCase{
		{
			"case 1",
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
	}

	for _, c := range cases {
		createTwoSumTestCase(t, c)
	}
}
