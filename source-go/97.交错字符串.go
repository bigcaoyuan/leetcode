/*
 * @lc app=leetcode.cn id=97 lang=golang
 *
 * [97] 交错字符串
 */
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	i, j := 0, 0
	for m := 0; m < len(s3); m++ {
		if i == len(s1) {
			return s3[m:] == s2[j:]
		}
		if j == len(s2) {
			return s3[m:] == s1[i:]
		}
		if s1[i] != s2[j] {
			if s3[m] == s1[i] {
				i++
				continue
			}
			if s3[m] == s2[j] {
				j++
				continue
			}
		} else {
			if s3[m] != s1[i] {
				return false
			}
			return isInterleave(s1[i+1:], s2[j:], s3[m+1:]) || isInterleave(s1[i:], s2[j+1:], s3[m+1:])
		}
		return false
	}
	return true
}

