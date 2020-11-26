package leetcode

import (
	"testing"
)

type LeeterCombinationsOfaPhoneNumberCase struct {
	name   string
	text   string
	result []string
}

func (c *LeeterCombinationsOfaPhoneNumberCase) ok(r []string) bool {
	return StringEquals(c.result, r, true)
}

func createLeeterCombinationsOfaPhoneNumberTestCase(t *testing.T, c *LeeterCombinationsOfaPhoneNumberCase) {
	t.Helper()
	result := letterCombinations(c.text)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestLetterCombinations(t *testing.T) {
	cases := []*LeeterCombinationsOfaPhoneNumberCase{
		{
			"case 1",
			"23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			"case 2",
			"",
			[]string{},
		},
		{
			"case 3",
			"7",
			[]string{"p", "q", "r", "s"},
		},
	}

	for _, c := range cases {
		createLeeterCombinationsOfaPhoneNumberTestCase(t, c)
	}

}
