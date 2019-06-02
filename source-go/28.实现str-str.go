/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现strStr()
 */
// 使用KMP算法
func strStr(s string, p string) int {
    if len(p) == 0{
        return 0 
    }
    pp := make([]int, len(p))
	pp[0] = 0
	for i := 1; i < len(p); i++ {
		j := pp[i-1]
		for {
			if p[i] == p[j] {
				pp[i] = j + 1
				break
			} else {
				if j == 0 { 
					pp[i] = 0
					break
				}
				j = pp[j-1]
			}
		}
	}

	i, j := 0, 0
	for i < len(s) && j < len(p) {
		if s[i] == p[j] {
			i++
			j++
		} else {
			if j == 0 {
				i++
				continue
			}
			j = pp[j-1]
		}
	}
	if j == len(p) {
		return i - j 
	}
	return -1
}


