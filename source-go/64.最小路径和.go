/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */
func minPathSum(grid [][]int) int {
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	for i := 1; i < len(grid[0]); i++ {
		grid[0][i] = grid[0][i-1] + grid[0][i]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			if grid[i-1][j] < grid[i][j-1] {
				grid[i][j] = grid[i][j] + grid[i-1][j]
			} else {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

