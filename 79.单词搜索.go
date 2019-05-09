/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 {
		return false
	}
	if len(board[0]) == 0 {
		return false
	}
	flag := make([][]int, len(board))
	for i := 0; i < len(flag); i++ {
		flag[i] = make([]int, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if doExist(&board, i, j, &word, 0, &flag) {
				return true
			}
		}
	}
	return false
}

// ["A","B","C","E"],
// ["S","F","C","S"],
// ["A","D","E","E"]]
// ABCCED

// ["C","A","A"],
// ["A","A","A"],
// ["B","C","D"]

// AAB

func doExist(board *[][]byte, i, j int, word *string, m int, flag *[][]int) bool {
	if i == len(*board) || i == -1 {
		return false
	}
	if j == len((*board)[0]) || j == -1 {
		return false
	}
	// 使用flag来防止回溯
	if (*flag)[i][j] == 1 {
		return false
	}
	if (*board)[i][j] == (*word)[m] {
		if m == len(*word)-1 {
			return true
		}
		(*flag)[i][j] = 1
		if doExist(board, i+1, j, word, m+1, flag) {
			return true
		}
		if doExist(board, i-1, j, word, m+1, flag) {
			return true
		}
		if doExist(board, i, j+1, word, m+1, flag) {
			return true
		}
		if doExist(board, i, j-1, word, m+1, flag) {
			return true
		}
		(*flag)[i][j] = 0
	}
	return false
}

