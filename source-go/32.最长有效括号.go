/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */
// (()())
func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	cnt := make([]int, len(s))
	max := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i-2 >= 0 {
					cnt[i] = cnt[i-2] + 2
				} else {
					cnt[i] = 2
				}
			} else if s[i-1] == ')' && (i-cnt[i-1]-1) >= 0 && s[i-cnt[i-1]-1] == '(' {
				if i-cnt[i-1]-2 < 0 {
					cnt[i] = cnt[i-1] + 2
				} else {
					cnt[i] = cnt[i-1] + 2 + cnt[i-cnt[i-1]-2]
				}
			}
			if cnt[i] > max {
				max = cnt[i]
			}
		}
	}
	return max
} 
