/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */
func isMatch(s string, p string) bool {
	return doisMatch(s, p, len(s)-1, len(p)-1)
}

// babcaacbccbbbbbabb
// bab*.b*b*a*aba*c

//
// c*
func doisMatch(s, p string, i, j int) bool {
	if i == -1 && j == -1 {
		return true
	} else if j == -1 {
		return false
	} else if i == -1 {
		if j%2 == 0 {
			return false
		}
		for j >= 0 {
			if p[j] != '*' {
				return false
			}
			j -= 2
		}
		return true
	}
	if s[i] == p[j] || p[j] == '.' {
		return doisMatch(s, p, i-1, j-1)
	} else if p[j] == '*' {
		if p[j-1] == s[i] || p[j-1] == '.' {
			// 匹配多次
			if doisMatch(s, p, i-1, j) {
				return true
			}
			// 只匹配一次
			if doisMatch(s, p, i-1, j-2) {
				return true
			}
		}
		return doisMatch(s, p, i, j-2)
	}
	return false
}

