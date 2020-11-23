package leetcode

import "testing"

type RomanToIntegerCase struct {
	name   string
	text   string
	result int
}

func (c *RomanToIntegerCase) ok(r int) bool {
	return c.result == r
}

func createRomanToIntegerTestCase(t *testing.T, c *RomanToIntegerCase) {
	t.Helper()
	result := romanToInt(c.text)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestRomanToInt(t *testing.T) {
	cases := []*RomanToIntegerCase{
		{
			"case 1",
			"III",
			3,
		},
		{
			"case 2",
			"IV",
			4,
		},
		{
			"case 3",
			"IX",
			9,
		},
		{
			"case 4",
			"LVIII",
			58,
		},
		{
			"case 5",
			"MCMXCIV",
			1994,
		},
	}

	for _, c := range cases {
		createRomanToIntegerTestCase(t, c)
	}

}
