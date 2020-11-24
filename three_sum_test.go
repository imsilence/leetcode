package leetcode

import (
	"sort"
	"testing"
)

type ThreeSumCase struct {
	name   string
	list   []int
	result [][]int
}

func (c *ThreeSumCase) ok(r [][]int) bool {
	if len(c.result) != len(r) {
		return false
	}
	sort.Slice(r, func(i, j int) bool {
		if r[i][0] == r[j][0] {
			return r[i][1] > r[j][1]
		}
		return r[i][0] > r[j][0]
	})

	equals := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		sort.Ints(a)
		sort.Ints(b)

		for i, v := range a {
			if v != b[i] {
				return false
			}
		}

		return true
	}

	for i, line := range r {
		if !equals(c.result[i], line) {
			return false
		}
	}

	return true
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
