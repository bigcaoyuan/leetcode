/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	isBlock := false
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			isBlock = true
			obstacleGrid[i][0] = -1
		} else {
			if isBlock {
				obstacleGrid[i][0] = -1
			} else {
				obstacleGrid[i][0] = 1
			}
		}
	}

	isBlock = false
	for i := 1; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 1 {
			isBlock = true
			obstacleGrid[0][i] = -1
		} else {
			if isBlock {
				obstacleGrid[0][i] = -1
			} else {
				obstacleGrid[0][i] = 1
			}
		}
	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = -1
				continue
			}
			m, n := obstacleGrid[i-1][j], obstacleGrid[i][j-1]
			if m == -1 {
				obstacleGrid[i][j] = n
			} else if n == -1 {
				obstacleGrid[i][j] = m
			} else {
				obstacleGrid[i][j] = m + n
			}
		}
	}
	if obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == -1 {
		return 0
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

