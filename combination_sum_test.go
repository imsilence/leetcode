package leetcode

import "testing"

type CombinationSumCase struct {
	name   string
	list   []int
	target int
	result [][]int
}

func (c *CombinationSumCase) ok(r [][]int) bool {
	return Int2Equals(c.result, r, true)
}

func createCombinationSumTestCase(t *testing.T, c *CombinationSumCase) {
	t.Helper()
	result := combinationSum(c.list, c.target)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestCombinationSum(t *testing.T) {
	cases := []*CombinationSumCase{
		{
			"case 1",
			[]int{2, 3, 6, 7},
			7,
			[][]int{
				{7},
				{2, 2, 3},
			},
		},
		{
			"case 2",
			[]int{2, 3, 5},
			8,
			[][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
	}

	for _, c := range cases {
		createCombinationSumTestCase(t, c)
	}

}
