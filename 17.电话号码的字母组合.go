/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */
func letterCombinations(digits string) []string {
	trans := map[rune][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	res := make([]string, 0)
	for _, v := range digits {
		if _, ok := trans[v]; !ok {
			continue
		}
		if len(res) == 0 {
			for i := 0; i < len(trans[v]); i++ {
				res = append(res, string(trans[v][i]))
			}
			continue
		} else {
			cnt := len(res)
			for i := 0; i < len(trans[v]); i++ {
				for j := 0; j < cnt; j++ {
					res = append(res, res[j]+string(trans[v][i]))
				}
			}
			res = res[cnt:]
		}
	}
	return res
}

