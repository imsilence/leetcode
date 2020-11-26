package leetcode

import (
	"testing"
)

type SubstringWithConcatenationOfAllWordsCase struct {
	name   string
	text   string
	list   []string
	result []int
}

func (c *SubstringWithConcatenationOfAllWordsCase) ok(r []int) bool {
	return IntEquals(c.result, r, true)
}

func createSubstringWithConcatenationOfAllWordsTestCase(t *testing.T, c *SubstringWithConcatenationOfAllWordsCase) {
	t.Helper()
	result := findSubstring(c.text, c.list)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestFindSubstring(t *testing.T) {
	cases := []*SubstringWithConcatenationOfAllWordsCase{
		{
			"case 1",
			"barfoothefoobarman",
			[]string{"foo", "bar"},
			[]int{0, 9},
		},
		{
			"case 2",
			"wordgoodgoodgoodbestword",
			[]string{"word", "good", "best", "word"},
			[]int{},
		},
	}

	for _, c := range cases {
		createSubstringWithConcatenationOfAllWordsTestCase(t, c)
	}
}
