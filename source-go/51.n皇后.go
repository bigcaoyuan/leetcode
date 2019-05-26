/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N皇后
 */

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	s := make([][]byte, n)
	for i := 0; i < n; i++ {
		s[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			s[i][j] = '.'
		}
	}
	doSolveNQueens(&s, 0, n, &res)
	return res
}

func doSolveNQueens(s *[][]byte, r, n int, res *[][]string) {
	if r == n {
		tmp := make([]string, n)
		for i := 0; i < n; i++ {
			tmp[i] = string((*s)[i])
		}
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < n; i++ {
		if check(s, r, i, n) {
			(*s)[r][i] = 'Q'
			doSolveNQueens(s, r+1, n, res)
			(*s)[r][i] = '.'
		}
	}
	return
}

func check(s *[][]byte, r, c, n int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (*s)[r][j] == 'Q' {
				return false
			}
			if (*s)[i][c] == 'Q' {
				return false
			}
			diffy := r - i
			diffx := c - j
			if diffy < 0 {
				diffy = -diffy
			}
			if diffx < 0 {
				diffx = -diffx
			}
			if diffy == diffx && (*s)[i][j] == 'Q' {
				return false
			}
		}
	}
	return true
}
