/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */
func mySqrt(x int) int {
	for i := 0; i <= x; i++ {
		temp := i * i
		if temp == x {
			return i
		}
		if temp > x {
			return i - 1
		}
	}
	return 0
}

