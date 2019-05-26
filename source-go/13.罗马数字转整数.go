/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */
func romanToInt(s string) int {
	mmap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 {
			if mmap[s[i]] < mmap[s[i+1]] {
				res -= mmap[s[i]]
				continue
			}
		}
		v := mmap[s[i]]
		res += v

	}
	return res
}

