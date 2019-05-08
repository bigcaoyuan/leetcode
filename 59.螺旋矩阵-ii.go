/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	j := 0
	i := 1
	for {
		if i > n*n {
			break
		}
		for m := j; m < n-j; m++ {
			res[j][m] = i
			i++
		}
		if i > n*n {
			break
		}
		for m := j + 1; m < n-j-1; m++ {
			res[m][n-j-1] = i
			i++
		}
		for m := n - j - 1; m >= j; m-- {
			res[n-j-1][m] = i
			i++
		}
		for m := n - j - 2; m >= j+1; m-- {
			res[m][j] = i
			i++
		}
		j++
	}
	return res
}

