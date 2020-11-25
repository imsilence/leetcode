package leetcode

import "testing"

type SearchInsertPositionCase struct {
	name   string
	list   []int
	target int
	result int
}

func (c *SearchInsertPositionCase) ok(r int) bool {
	return c.result == r
}

func createSearchInsertPositionTestCase(t *testing.T, c *SearchInsertPositionCase) {
	t.Helper()
	result := searchInsert(c.list, c.target)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestSearchInsert(t *testing.T) {
	cases := []*SearchInsertPositionCase{
		{
			"case 1",
			[]int{1, 3, 5, 6},
			5,
			2,
		},
		{
			"case 2",
			[]int{1, 3, 5, 6},
			2,
			1,
		},
		{
			"case 3",
			[]int{1, 3, 5, 6},
			7,
			4,
		},
		{
			"case 4",
			[]int{1, 3, 5, 6},
			0,
			0,
		},
	}

	for _, c := range cases {
		createSearchInsertPositionTestCase(t, c)
	}

}
