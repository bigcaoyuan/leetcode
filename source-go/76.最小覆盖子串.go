/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */
func minWindow(s string, t string) string {
	fre := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		fre[t[i]]++
	}
	l, r, tmp, res := 0, 1, make(map[byte]int), ""
	tmp[s[0]]++
	for {
		ok := checkmp(tmp, fre)
		if !ok {
			if r == len(s) {
				break
			}
			tmp[s[r]]++
			r++
		} else {
			if len(res) == 0 || r-l < len(res) {
				res = s[l:r]
			}
			tmp[s[l]]--
			l++
		}
	}
	return res
}

func checkmp(a, b map[byte]int) bool {
	for k, v := range b {
		if v > a[k] {
			return false
		}
	}
	return true
}

