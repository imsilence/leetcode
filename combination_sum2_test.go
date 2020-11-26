package leetcode

import "testing"

type CombinationSum2Case struct {
	name   string
	list   []int
	target int
	result [][]int
}

func (c *CombinationSum2Case) ok(r [][]int) bool {
	return Int2Equals(c.result, r, true)
}

func createCombinationSum2TestCase(t *testing.T, c *CombinationSum2Case) {
	t.Helper()
	result := combinationSum2(c.list, c.target)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestCombinationSum2(t *testing.T) {
	cases := []*CombinationSum2Case{
		{
			"case 1",
			[]int{2, 3, 6, 7},
			7,
			[][]int{
				{7},
			},
		},
		{
			"case 2",
			[]int{2, 3, 5},
			8,
			[][]int{
				{3, 5},
			},
		},
		{
			"case 3",
			[]int{10, 1, 2, 7, 6, 1, 5},
			8,
			[][]int{
				{1, 7},
				{1, 2, 5},
				{2, 6},
				{1, 1, 6},
			},
		},
		{
			"case 4",
			[]int{2, 5, 2, 1, 2},
			5,
			[][]int{
				{1, 2, 2},
				{5},
			},
		},
		{
			"case 5",
			[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			27,
			nil,
		},
	}

	for _, c := range cases {
		createCombinationSum2TestCase(t, c)
	}

}
