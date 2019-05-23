/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */
func myAtoi(str string) int {
	num := 0
	hasFlag := false
	negivate := 1
	hasTrim := false
	MIN_INT := -2147483648
	MAX_INT := 2147483647
	for i := 0; i < len(str); i++ {
		// 去除左侧空格
		if !hasTrim && str[i] == ' ' {
			continue
		} else {
			hasTrim = true
		}
		// 起始为非法字符如 w
		if str[i] < '0' || str[i] > '9' {
			if hasFlag {
				return num * negivate
			} else {
				if str[i] == '-' {
					hasFlag = true
					negivate = -1
					continue
				} else if str[i] == '+' {
					hasFlag = true
					continue
				} else {
					return 0
				}
			}
		}
		hasFlag = true
		num = num*10 + (int(str[i]) - 48)
		if num*negivate > MAX_INT {
			return MAX_INT
		} else if num*negivate < MIN_INT {
			return MIN_INT
		}
	}
	return num * negivate
}

