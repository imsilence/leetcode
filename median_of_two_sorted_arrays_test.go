package leetcode

import (
	"math"
	"testing"
)

type MedianOfTwoSortedArraysCase struct {
	name   string
	left   []int
	right  []int
	result float64
}

func (c *MedianOfTwoSortedArraysCase) ok(r float64) bool {
	return math.Abs(c.result-r) < 0.001
}

func createMedianOfTwoSortedArraysTestCase(t *testing.T, c *MedianOfTwoSortedArraysCase) {
	result := findMedianSortedArrays(c.left, c.right)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	cases := []*MedianOfTwoSortedArraysCase{
		{
			"case 1",
			[]int{1, 3},
			[]int{2},
			2,
		},
		{
			"case 2",
			[]int{1, 2},
			[]int{3, 4},
			2.5,
		},
		{
			"case 3",
			[]int{0, 0},
			[]int{0, 0},
			0,
		},
		{
			"case 4",
			[]int{1},
			[]int{},
			1,
		},
		{
			"case 5",
			[]int{},
			[]int{2},
			2,
		},
	}
	for _, c := range cases {
		createMedianOfTwoSortedArraysTestCase(t, c)
	}
}
