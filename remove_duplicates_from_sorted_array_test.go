package leetcode

import "testing"

type RemoveDuplicatesFromSortedArrayCase struct {
	name   string
	list   []int
	result int
}

func (c *RemoveDuplicatesFromSortedArrayCase) ok(r int) bool {
	return c.result == r
}

func createRemoveDuplicatesFromSortedArrayTestCase(t *testing.T, c *RemoveDuplicatesFromSortedArrayCase) {
	t.Helper()
	result := removeDuplicates(c.list)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	cases := []*RemoveDuplicatesFromSortedArrayCase{
		{
			"case 1",
			[]int{1, 1, 2},
			2,
		},
		{
			"case 2",
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
		},
	}

	for _, c := range cases {
		createRemoveDuplicatesFromSortedArrayTestCase(t, c)
	}

}
