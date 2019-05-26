/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */
func lengthOfLastWord(s string) int {
	cnt := 0
	trim := true
	for i := len(s) - 1; i >= 0; i-- {
		if trim && s[i] == 32 {
			continue
		}
		if s[i] != 32 {
			cnt++
			trim = false
			continue
		}
		break
	}
	return cnt
}

