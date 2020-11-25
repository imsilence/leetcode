package leetcode

import (
	"testing"
)

type CountAndSayCase struct {
	name   string
	num    int
	result string
}

func (c *CountAndSayCase) ok(r string) bool {
	return c.result == r
}

func createCountAndSayTestCase(t *testing.T, c *CountAndSayCase) {
	t.Helper()
	result := countAndSay(c.num)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestCountAndSay(t *testing.T) {
	cases := []*CountAndSayCase{
		{
			"case 1",
			1,
			"1",
		},
		{
			"case 2",
			2,
			"11",
		},
		{
			"case 3",
			3,
			"21",
		},
		{
			"case 4",
			4,
			"1211",
		},
		{
			"case 5",
			5,
			"111221",
		},
	}

	for _, c := range cases {
		createCountAndSayTestCase(t, c)
	}

}
