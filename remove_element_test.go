package leetcode

import "testing"

type RemoveElementCase struct {
	name   string
	list   []int
	target int
	result []int
}

func (c *RemoveElementCase) ok(r int) bool {
	if len(c.result) != r {
		return false
	}
	for i, v := range c.result {
		if v != c.list[i] {
			return false
		}
	}
	return true
}

func createRemoveElementTestCase(t *testing.T, c *RemoveElementCase) {
	t.Helper()
	result := removeElement(c.list, c.target)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestRemoveElement(t *testing.T) {
	cases := []*RemoveElementCase{
		{
			"case 1",
			[]int{3, 2, 2, 3},
			3,
			[]int{2, 2},
		},
		{
			"case 2",
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			[]int{0, 1, 3, 0, 4},
		},
	}

	for _, c := range cases {
		createRemoveElementTestCase(t, c)
	}

}
