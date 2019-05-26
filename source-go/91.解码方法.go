/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */
func numDecodings(s string) int {
	// 48 49 50 51 52 53 54 55 56 57
	// 0  1  2  3  4  5  6  7  8  9
	// 120
	if s[0] == '0' {
		return 0
	}
	cnt := make([]int, len(s)+1)
	cnt[0] = 1
	cnt[1] = 1
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cnt[i+1] = cnt[i-1]
			} else {
				return 0
			}
		} else if s[i-1] == '1' {
			cnt[i+1] = cnt[i] + cnt[i-1]
		} else if s[i-1] == '2' {
			if s[i] < '7' {
				cnt[i+1] = cnt[i] + cnt[i-1]
			} else {
				cnt[i+1] = cnt[i]
			}
		} else {
			cnt[i+1] = cnt[i]
		}
	}
	return cnt[len(s)]
} 
