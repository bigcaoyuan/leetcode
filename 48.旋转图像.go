/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */
func rotate(matrix [][]int) {
	n := len(matrix)
	if n <= 1 {
		return
	}
	i, j := 0, 0
	for {
		matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] = matrix[n-1-j][i], matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j]
		j++
		if j == n-1-i {
			i++
			j = i
			if i == n/2 {
				break
			}

		}
	}
	return
}

