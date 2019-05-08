/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	gen("", n, 0, 0, &res)
	return res
}

func gen(str string, n, left, right int, res *[]string) {
	if right == n {
		*res = append(*res, str)
		return
	}
	if left < n {
		gen(str+"(", n, left+1, right, res)
	}
	if right < left {
		gen(str+")", n, left, right+1, res)
	}
}

