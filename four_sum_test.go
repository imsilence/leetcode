package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

type FourSumCase struct {
	name   string
	nums   []int
	target int
	result [][]int
}

func (c *FourSumCase) ok(r [][]int) bool {
	if len(c.result) != len(r) {
		return false
	}
	for _, v := range r {
		sort.Ints(v)
	}
	sort.Slice(r, func(i, j int) bool {
		for k, v := range r[i] {
			if v < r[j][k] {
				return true
			} else if v > r[j][k] {
				return false
			}
		}
		return false
	})

	fmt.Println(r)
	equals := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}

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

func createFourSumTestCase(t *testing.T, c *FourSumCase) {
	t.Helper()
	result := fourSum(c.nums, c.target)
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
