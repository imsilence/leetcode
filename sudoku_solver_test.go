package leetcode

import (
	"testing"
)

type SudokuSolverCase struct {
	name   string
	list   [][]byte
	result [][]byte
}

func (c *SudokuSolverCase) ok() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if c.list[i][j] != c.result[i][j] {
				return false
			}
		}
	}
	return true
}

func createSudokuSolverTestCase(t *testing.T, c *SudokuSolverCase) {
	t.Helper()
	solveSudoku(c.list)
	if !c.ok() {
		t.Error(c.name, c)
	}
}

func TestSolveSudoku(t *testing.T) {
	cases := []*SudokuSolverCase{
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
			[][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
			},
		},
	}

	for _, c := range cases {
		createSudokuSolverTestCase(t, c)
	}
}
