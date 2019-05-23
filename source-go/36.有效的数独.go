/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */
func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	index := make([]map[byte]bool, 9)
	for i := 0; i < len(index); i++ {
		index[i] = make(map[byte]bool)
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val == '.' {
				continue
			}
			if _, ok := rows[i][val]; ok {
				return false
			}
			if _, ok := cols[j][val]; ok {
				return false
			}
			if _, ok := index[i/3*3+j/3][val]; ok {
				return false
			}
			rows[i][val] = true
			cols[j][val] = true
			index[i/3*3+j/3][val] = true
		}
	}
	return true
}

