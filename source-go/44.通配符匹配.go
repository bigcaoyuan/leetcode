/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 */
// 动态规划
func isMatch(s string, p string) bool {
	bools := make([][]bool, len(p)+1)
	for i := 0; i < len(bools); i++ {
		bools[i] = make([]bool, len(s)+1)
	}
	for i := 1; i < len(bools[0]); i++ {
		bools[0][i] = false
	}
	bools[0][0] = true
	for i := 1; i <= len(p); i++ {
		if p[i-1] == '?' {
			bools[i][0] = false
			for j := 1; j < len(bools[0]); j++ {
				bools[i][j] = bools[i-1][j-1]
			}
		} else if p[i-1] == '*' {
			bools[i][0] = bools[i-1][0]
			for j := 1; j < len(bools[0]); j++ {
				bools[i][j] = bools[i-1][j-1] || bools[i-1][j] || bools[i][j-1]
			}
		} else {
			bools[i][0] = false
			for j := 1; j < len(bools[0]); j++ {
				bools[i][j] = bools[i-1][j-1] && p[i-1] == s[j-1]
			}
		}
	}
	return bools[len(p)][len(s)]
}

// 回溯
func isMatch2(s string, p string) bool {
	i, j, start, k := 0, 0, -1, 0

	for i < len(s) {
		if j < len(p) && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
			continue
		}
		// 这一次跳过*
		if j < len(p) && p[j] == '*' {
			k = j
			start = i
			j++
			continue
		}
		if start != -1 {
			i = start + 1
			start++
			j = k + 1
			continue
		}
		return false
	}
	for ; j < len(p); j++ {
		if p[j] != '*' {
			return false
		}
	}
	return i == len(s) && j == len(p)
}

func isMatch1(s string, p string) bool {
	for i := 0; i < len(p)-1; i++ {
		if p[i] == '*' && p[i] == p[i+1] {
			p = p[:i] + p[i+1:]
			i--
		}
	}
	return doIsMatch(s, p, 0, 0)
}

// adceb
// *a*b
func doIsMatch(s string, p string, i, j int) bool {
	if i == len(s) || j == len(p) {
		if i == len(s) && j == len(p) {
			return true
		}
		if i == len(s) && p[j] == '*' {
			return doIsMatch(s, p, i, j+1)
		}
		return false
	}
	if s[i] == p[j] || p[j] == '?' {
		return doIsMatch(s, p, i+1, j+1)
	} else if p[j] == '*' {
		return doIsMatch(s, p, i+1, j) || doIsMatch(s, p, i, j+1)
	} else {
		return false
	}
}

