/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	cnt := numRows*2 - 2
	str := ""
	for i := 0; i < numRows; i++ {
		j := 0
		for {
			cc := j * cnt
			if cc >= len(s) {
				break
			}
			if cc+i >= len(s) {
				break
			}
			str += string(s[cc+i])
			index := cnt - i
			if i != 0 && index != i {
				if cc+index >= len(s) {
					break
				}
				str += string(s[cc+index])
			}
			j++
		}
	}
	return str
}

