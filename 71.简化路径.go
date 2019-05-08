import "strings"

/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */
func simplifyPath(path string) string {
	items := strings.Split(path, "/")
	res := make([]string, 0)
	for i := 0; i < len(items); i++ {
		if items[i] == "." || len(items[i]) == 0 {
			continue
		}
		if items[i] == ".." {
			if len(res) != 0 {
				res = res[:len(res)-1]
			}
			continue
		}
		res = append(res, items[i])
	}
	str := strings.Join(res, "/")
	return "/" + str
}

