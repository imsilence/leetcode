package leetcode

import (
	"testing"
)

type ZigzagConversionCase struct {
	name   string
	text   string
	rows   int
	result string
}

func (c *ZigzagConversionCase) ok(r string) bool {
	return c.result == r
}

func createZigzagConversionTestCase(t *testing.T, c *ZigzagConversionCase) {
	t.Helper()
	result := convert(c.text, c.rows)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestConvert(t *testing.T) {
	cases := []*ZigzagConversionCase{
		{
			"case 1",
			"LEETCODEISHIRING",
			3,
			"LCIRETOESIIGEDHN",
		},
		{
			"case 2",
			"LEETCODEISHIRING",
			4,
			"LDREOEIIECIHNTSG",
		},
		{
			"case 3",
			"AB",
			1,
			"AB",
		},
	}

	for _, c := range cases {
		createZigzagConversionTestCase(t, c)
	}
}
