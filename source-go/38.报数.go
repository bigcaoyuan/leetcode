import "fmt"

/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 报数
 */

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := "1"
	for i := 2; i <= n; i++ {
		temp := ""
		lastVal := str[0]
		lastCnt := 1
		for j := 1; j < len(str); j++ {
			if str[j] == lastVal {
				lastCnt++
			} else {
				temp += fmt.Sprintf("%d", lastCnt)
				temp += string(lastVal)
				lastVal = str[j]
				lastCnt = 1
			}
		}
		temp += fmt.Sprintf("%d", lastCnt)
		temp += string(lastVal)
		str = temp
	}
	return str
}

