/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x%10 == 0 && x != 0 {
		return false
	}
	res := 0
	for x > res {
		res = res*10 + x%10
		x = x / 10
	}

	if res == x {
		return true
	}
	if res/10 == x {
		return true
	}

	return false
}

