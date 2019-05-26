/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

//  ADOBECODEBANC
// ABC
// BANC
func minWindow(s string, t string) string {
	if len(t) == 0 {
		return t
	}
	fre := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		fre[t[i]]++
	}
	tmp := make(map[byte]int)
	l, r, lent := 0, 0, len(t)
	str := ""
	for {
		if r == len(s)+1 {
			break
		}
		ok := checkmp(tmp, fre)
		if ok {
			if r-l+1 == lent {
				return s[l:r]
			}
			if len(str) == 0 || r-l+1 < len(str) {
				str = s[l:r]
			}
			tmp[s[l]]--
			l++
		} else {
			if r < len(s) {
				tmp[s[r]]++
				r++
			} else {
				break
			}
		}
	}
	return str
}

func checkmp(a, b map[byte]int) bool {
	for k, v := range b {
		if vv, ok := a[k]; !ok {
			return false
		} else if vv < v {
			return false
		}
	}
	return true
}

