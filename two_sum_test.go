package leetcode

import (
	"sort"
	"testing"
)

type TwoSumCase struct {
	name   string
	list   []int
	target int
	result []int
}

func (c *TwoSumCase) ok(r []int) bool {

	if len(c.result) != len(r) {
		return false
	}

	sort.Ints(c.result)
	sort.Ints(r)

	for i, v := range c.result {
		if v != r[i] {
			return false
		}
	}
	return true
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
