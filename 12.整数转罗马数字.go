/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */
func intToRoman(num int) string {
	strArr := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	intArr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	str := ""
	for i, val := range intArr {
		cnt := num / val
		for j := 0; j < cnt; j++ {
			str = str + strArr[i]
		}
		num = num - val*cnt
	}
	return str
}

