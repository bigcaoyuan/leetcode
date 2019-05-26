/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 == 10 {
			digits[i] = 0
		} else {
			digits[i] = digits[i] + 1
			break
		}
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

