/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */
func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	if i < j {
		for m := 0; m < j-i; m++ {
			a = "0" + a
		}
		i = j
	} else if j < i {
		for m := 0; m < i-j; m++ {
			b = "0" + b
		}
		j = i
	}
	var res string
	add := byte(0)
	for {
		if i < 0 {
			break
		}
		// 0 = 48  1= 49
		sum := a[i] + b[j] + add
		val := sum % byte(2)
		if sum >= 98 {
			add = sum - val - 97
		} else {
			add = 0
		}
		if val == byte(1) {
			res = "1" + res
		} else {
			res = "0" + res
		}
		i--
		j--
	}
	if add != 0 {
		res = "1" + res
	}

	return res
}

