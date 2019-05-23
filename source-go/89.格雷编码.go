/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	cnt := 1
	for i := 1; i < n; i++ {
		cnt = cnt << 1
	}
	res := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		res[i] = i ^ (i >> 1)
	}
	return res
}

