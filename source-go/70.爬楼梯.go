/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */
func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	last := 2
	lastlast := 1
	cnt := 0
	for i := 3; i <= n; i++ {
		cnt = lastlast + last
		lastlast = last
		last = cnt
	}

	return cnt
}
