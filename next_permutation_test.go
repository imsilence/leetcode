package leetcode

import "testing"

type NextPermutationCase struct {
	name   string
	list   []int
	result []int
}

func (c *NextPermutationCase) ok(r []int) bool {
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

func createNextPermutationCase(t *testing.T, c *NextPermutationCase) {
	t.Helper()
	nextPermutation(c.list)
	if !c.ok(c.list) {
		t.Error(c.name, c, c.list)
	}
}

func TestNextPermutation(t *testing.T) {
	cases := []*NextPermutationCase{
		{
			"case 1",
			[]int{1, 1, 5},
			[]int{1, 5, 1},
		},
		{
			"case 2",
			[]int{1},
			[]int{1},
		},
		{
			"case 3",
			[]int{1, 2, 3},
			[]int{1, 3, 2},
		},
		{
			"case 4",
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
		{
			"case 5",
			[]int{1, 3, 2},
			[]int{2, 1, 3},
		},
		{
			"case 6",
			[]int{2, 3, 1},
			[]int{3, 1, 2},
		},
	}

	for _, c := range cases {
		createNextPermutationCase(t, c)
	}

}
