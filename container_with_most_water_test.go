package leetcode

import "testing"

type ContainerWithMostWaterCase struct {
	name   string
	nums   []int
	result int
}

func (c *ContainerWithMostWaterCase) ok(r int) bool {
	return c.result == r
}

func createContainerWithMostWaterTestCase(t *testing.T, c *ContainerWithMostWaterCase) {
	t.Helper()
	result := maxArea(c.nums)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestMaxArea(t *testing.T) {
	cases := []*ContainerWithMostWaterCase{
		{
			"case 1",
			[]int{1, 1},
			1,
		},
		{
			"case 2",
			[]int{4, 3, 2, 1, 4},
			16,
		},
		{
			"case 3",
			[]int{1, 2, 1},
			2,
		},
	}

	for _, c := range cases {
		createContainerWithMostWaterTestCase(t, c)
	}

}
