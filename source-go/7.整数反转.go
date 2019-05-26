/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
func reverse(x int) int {
	res := 0
	Infl := 214748364
	isnNegative := false
	if x < 0 {
		x = -x
		isnNegative = true
	}
	for x != 0 {
		n := x % 10
		if res > Infl || (res == Infl && n > 7) {
			return 0
		}
		res = res*10 + n
		x = x / 10
	}
	if isnNegative {
		res = -res
	}
	return res
}

