/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
func longestPalindrome(s string) string {
	// 动态规划
	maxsum := 0
	maxStr := ""
	for i := 0; i < len(s); i++ {
		for j := 0; ; j++ {
			if i-j < 0 || i+j >= len(s) {
				break
			}
			if s[i-j] == s[i+j] {
				if (j*2 + 1) > maxsum {
					maxsum = j*2 + 1
					maxStr = s[i-j : i+j+1]
				}
			} else {
				break
			}
		}
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			for j := 0; ; j++ {
				if i-j < 0 || i+j+1 >= len(s) {
					break
				}
				if s[i-j] == s[i+1+j] {
					if (j*2 + 2) > maxsum {
						maxsum = j*2 + 2
						maxStr = s[i-j : i+j+2]
					}
				} else {
					break
				}
			}
		}
	}
	return maxStr
}

