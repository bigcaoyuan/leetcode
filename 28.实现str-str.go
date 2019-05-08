/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现strStr()
 */
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	j := 0
	i := 0
	for i = 0; i < len(haystack); i++ {
		if j == len(needle) {
			return i - len(needle)
		}
		if haystack[i] == needle[j] {
			j++
		} else {
			i -= j
			j = 0
		}
	}
	if j == len(needle) {
		return i - len(needle)
	}
	return -1
}

