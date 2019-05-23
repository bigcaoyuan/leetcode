/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	start, end := 0, m*n-1
	for {
		if start > end {
			return false
		}
		mid := (start + end) / 2
		i := mid / n
		j := mid % n
		if target == matrix[i][j] {
			return true
		} else if target > matrix[i][j] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

