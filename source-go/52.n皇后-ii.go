/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 */
func totalNQueens(n int) int {
	cnt := 0
	s := make([][]byte, n)
	for i := 0; i < n; i++ {
		s[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			s[i][j] = '.'
		}
	}
	doSolveNQueens(&s, 0, n, &cnt)
	return cnt
}
func doSolveNQueens(s *[][]byte, r, n int, cnt *int) {
	if r == n {
		(*cnt)++
		return
	}
	for i := 0; i < n; i++ {
		if check(s, r, i, n) {
			(*s)[r][i] = 'Q'
			doSolveNQueens(s, r+1, n, cnt)
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
