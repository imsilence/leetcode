package leetcode

import "testing"

type FindFirstAndLastPositionOfElementInSortedArrayCase struct {
	name   string
	list   []int
	target int
	result []int
}

func (c *FindFirstAndLastPositionOfElementInSortedArrayCase) ok(r []int) bool {
	if len(c.result) != len(r) {
		return false
	}
	for i, v := range c.result {
		if v != r[i] {
			return false
		}
	}
	return true
}

func createFindFirstAndLastPositionOfElementInSortedArrayTestCase(t *testing.T, c *FindFirstAndLastPositionOfElementInSortedArrayCase) {
	t.Helper()
	result := searchRange(c.list, c.target)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestSearchRange(t *testing.T) {
	cases := []*FindFirstAndLastPositionOfElementInSortedArrayCase{
		{
			"case 1",
			[]int{5, 7, 7, 8, 8, 10},
			8,
			[]int{3, 4},
		},
		{
			"case 2",
			[]int{5, 7, 7, 8, 8, 10},
			6,
			[]int{-1, -1},
		},
		{
			"case 3",
			[]int{1},
			1,
			[]int{0, 0},
		},
	}

	for _, c := range cases {
		createFindFirstAndLastPositionOfElementInSortedArrayTestCase(t, c)
	}

}
