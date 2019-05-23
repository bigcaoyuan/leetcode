/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */
func multiply(num1 string, num2 string) string {
	res := make([]byte, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			val := (num1[i] - '0') * (num2[j] - '0')
			index := i + j + 1
			res[index] += val
		}
		add := byte(0)
		for i := len(res) - 1; i >= 0; i-- {
			if res[i]+add >= 10 {
				odd := (res[i] + add) % 10
				add = (res[i] + add) / 10
				res[i] = +odd
			} else {
				res[i] = res[i] + add
				add = byte(0)
			}
		}
	}

	add := byte(0)
	for i := len(res) - 1; i >= 0; i-- {
		if res[i]+add >= 10 {
			odd := (res[i] + add) % 10
			add = (res[i] + add) / 10
			res[i] = '0' + odd
		} else {
			res[i] = '0' + res[i] + add
			add = byte(0)
		}
	}
	zero := false
	str := ""
	for i := 0; i < len(res); i++ {
		if !zero && res[i] != '0' {
			zero = true
		}
		if zero {
			str = str + string(res[i])
		}
	}
	if str == "" {
		return "0"
	}
	return str
}


