/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */
func minDistance(word1 string, word2 string) int {
	mp := make([][]int, len(word1)+1)
	for i := 0; i < len(mp); i++ {
		mp[i] = make([]int, len(word2)+1)
		mp[i][0] = i
	}
	for i := 0; i < len(mp[0]); i++ {
		mp[0][i] = i
	}

	for i := 1; i < len(mp); i++ {
		for j := 1; j < len(mp[0]); j++ {
			if word1[i-1] == word2[j-1] {
				mp[i][j] = mp[i-1][j-1]
			} else {
				mp[i][j] = 1 + min(mp[i-1][j-1], min(mp[i-1][j], mp[i][j-1]))
			}
		}
	}
	return mp[len(word1)][len(word2)]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

