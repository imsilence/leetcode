package leetcode

import "testing"

type IntegerToRomanCase struct {
	name   string
	num    int
	result string
}

func (c *IntegerToRomanCase) ok(r string) bool {
	return c.result == r
}

func createIntegerToRomanTestCase(t *testing.T, c *IntegerToRomanCase) {
	t.Helper()
	result := intToRoman(c.num)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestIntToRoman(t *testing.T) {
	cases := []*IntegerToRomanCase{
		{
			"case 1",
			3,
			"III",
		},
		{
			"case 2",
			4,
			"IV",
		},
		{
			"case 3",
			9,
			"IX",
		},
		{
			"case 4",
			58,
			"LVIII",
		},
		{
			"case 5",
			1994,
			"MCMXCIV",
		},
	}

	for _, c := range cases {
		createIntegerToRomanTestCase(t, c)
	}

}
