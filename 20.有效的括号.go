/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
func isValid(s string) bool {
	open := map[rune]rune{
		'[': ']',
		'{': '}',
		'(': ')',
	}
	close := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
	stack := make([]rune, 0)
	for _, val := range s {
		if _, ok := open[val]; ok {
			stack = append(stack, val)
			continue
		}
		if corr, ok := close[val]; ok {
			if len(stack) == 0 {
				return false
			}
			if corr == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
			continue
		}
		return false
	}
	return len(stack) == 0
}

