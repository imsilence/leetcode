package leetcode

import (
	"testing"
)

type StringToIntegerAtoiCase struct {
	name   string
	text   string
	result int
}

func (c *StringToIntegerAtoiCase) ok(r int) bool {
	return c.result == r
}

func createStringToIntegerAtoiTestCase(t *testing.T, c *StringToIntegerAtoiCase) {
	t.Helper()
	result := myAtoi(c.text)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestMyAtoi(t *testing.T) {
	cases := []*StringToIntegerAtoiCase{
		// {
		// 	"case 1",
		// 	"42",
		// 	42,
		// },
		// {
		// 	"case 2",
		// 	"   -42",
		// 	-42,
		// },
		// {
		// 	"case 3",
		// 	"4193 with words",
		// 	4193,
		// },
		// {
		// 	"case 4",
		// 	"words and 987",
		// 	0,
		// },
		// {
		// 	"case 5",
		// 	"-91283472332",
		// 	-2147483648,
		// },
		// {
		// 	"case 6",
		// 	"+1",
		// 	1,
		// },
		// {
		// 	"case 7",
		// 	"9223372036854775808",
		// 	2147483647,
		// },
		// {
		// 	"case 8",
		// 	"-9223372036854775808",
		// 	-2147483648,
		// },
		{
			"case 9",
			"18446744073709551617",
			2147483647,
		},
	}

	for _, c := range cases {
		createStringToIntegerAtoiTestCase(t, c)
	}
}
