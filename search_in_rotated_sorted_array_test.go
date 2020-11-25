package leetcode

import "testing"

type SearchInRotatedSortedArrayCase struct {
	name   string
	list   []int
	target int
	result int
}

func (c *SearchInRotatedSortedArrayCase) ok(r int) bool {
	return c.result == r
}

func createSearchInRotatedSortedArrayTestCase(t *testing.T, c *SearchInRotatedSortedArrayCase) {
	t.Helper()
	result := search(c.list, c.target)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestSearch(t *testing.T) {
	cases := []*SearchInRotatedSortedArrayCase{
		{
			"case 1",
			[]int{4, 5, 6, 7, 0, 1, 2},
			0,
			4,
		},
		{
			"case 2",
			[]int{4, 5, 6, 7, 0, 1, 2},
			3,
			-1,
		},
		{
			"case 3",
			[]int{1},
			0,
			-1,
		},
	}

	for _, c := range cases {
		createSearchInRotatedSortedArrayTestCase(t, c)
	}

}
