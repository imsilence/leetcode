package leetcode

/*
37. 解数独
编写一个程序，通过填充空格来解决数独问题。

一个数独的解法需遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
空白格用 '.' 表示。



一个数独。



答案被标成红色。

提示：

给定的数独序列只包含数字 1-9 和字符 '.' 。
你可以假设给定的数独只有唯一解。
给定数独永远是 9x9 形式的。
*/

func check(board [][]byte, i, j int) bool {
	if board[i][j] == '.' {
		return true
	}

	checkLine := func(i int) bool {
		exists := map[byte]int{}
		for j := 0; j < 9; j++ {
			ch := board[i][j]
			if ch == '.' {
				continue
			}
			exists[ch]++
			if exists[ch] == 2 {
				return false
			}
		}
		return true
	}

	checkColumn := func(j int) bool {
		exists := map[byte]int{}
		for i := 0; i < 9; i++ {
			ch := board[i][j]
			if ch == '.' {
				continue
			}
			exists[ch]++
			if exists[ch] == 2 {
				return false
			}
		}
		return true
	}

	checkBlock := func(i, j int) bool {
		exists := map[byte]int{}
		x, y := (i/3)*3, (j/3)*3
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				ch := board[x+i][y+j]
				if ch == '.' {
					continue
				}
				exists[ch]++
				if exists[ch] == 2 {
					return false
				}
			}
		}
		return true
	}
	return checkLine(i) && checkColumn(j) && checkBlock(i, j)
}

func solve(board [][]byte) bool {
	text := "123456789"

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}
			for _, ch := range text {
				board[i][j] = byte(ch)
				if check(board, i, j) && solve(board) {
					return true
				}
			}
			board[i][j] = '.'
			return false
		}
	}
	return true
}

func solveSudoku(board [][]byte) {
	solve(board)
}
