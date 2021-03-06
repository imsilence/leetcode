package leetcode

import (
	"testing"
)

type ValidSudokuCase struct {
	name   string
	list   [][]byte
	result bool
}

func (c *ValidSudokuCase) ok(r bool) bool {
	return c.result == r
}

func createValidSudokuTestCase(t *testing.T, c *ValidSudokuCase) {
	t.Helper()
	result := isValidSudoku(c.list)
	if !c.ok(result) {
		t.Error(c.name, c, result)
	}
}

func TestIsValidSudoku(t *testing.T) {
	cases := []*ValidSudokuCase{
		{
			"case 1",
			[][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			true,
		},
		{
			"case 2",
			[][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			false,
		},
		{
			"case 3",
			[][]byte{
				{'.', '4', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '1', '.', '.', '7', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '3', '.', '.', '.', '6', '.'},
				{'.', '.', '.', '.', '.', '6', '.', '9', '.'},
				{'.', '.', '.', '.', '1', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
				{'.', '.', '.', '8', '.', '.', '.', '.', '.'},
			},
			false,
		},
	}

	for _, c := range cases {
		createValidSudokuTestCase(t, c)
	}
}
