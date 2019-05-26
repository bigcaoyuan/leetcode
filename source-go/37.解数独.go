/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */
func solveSudoku(board [][]byte) {
	items := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	doSolveSudoku(&items, &board, 0, 0)
}

func doSolveSudoku(items *[]byte, board *[][]byte, r, c int) bool {
	if c == 9 {
		r = r + 1
		c = 0
	}
	if r == 9 {
		return true
	}
	if (*board)[r][c] != '.' {
		return doSolveSudoku(items, board, r, c+1)
	}

	for m := 0; m < 9; m++ {
		if check(board, r, c, (*items)[m]) {
			(*board)[r][c] = (*items)[m]
			if doSolveSudoku(items, board, r, c+1) {
				return true
			}
		}
	}
	(*board)[r][c] = '.'
	return false
}

func check(board *[][]byte, r, c int, val byte) bool {
	for i := 0; i < 9; i++ {
		if (*board)[i][c] == val {
			return false
		}
		if (*board)[r][i] == val {
			return false
		}
	}
	r = r / 3
	c = c / 3
	for i := r * 3; i < (r+1)*3; i++ {
		for j := c * 3; j < (c+1)*3; j++ {
			if (*board)[i][j] == val {
				return false
			}
		}
	}
	return true
}

