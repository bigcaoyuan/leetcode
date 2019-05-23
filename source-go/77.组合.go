/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	if n == 0 {
		return res
	}
	items := make([]int, n)
	for i := 1; i <= n; i++ {
		items[i-1] = i
	}
	for i := 0; i < len(items); i++ {
		doCombine(items, []int{items[i]}, i+1, k, &res)
	}
	return res
}
func doCombine(items []int, local []int, index, k int, res *[][]int) {
	if len(local) == k {
		*res = append(*res, local)
		return
	}
	for i := index; i < len(items); i++ {
		temp := append([]int{}, local...)
		temp = append(temp, items[i])
		doCombine(items, temp, i+1, k, res)
	}
}

//全组合排列 搞混了
func combine1(n int, k int) [][]int {
	res := make([][]int, 0)
	if n == 0 {
		return res
	}
	items := make([]int, n)
	for i := 1; i <= n; i++ {
		items[i-1] = i
	}
	doCombine1(items, []int{}, k, &res)
	return res
}

func doCombine1(items []int, local []int, k int, res *[][]int) {
	if len(local) == k {
		*res = append(*res, local)
		return
	}
	for i := 0; i < len(items); i++ {
		temp := append([]int{}, local...)
		temp = append(temp, items[i])

		temp1 := append([]int{}, items[:i]...)
		temp1 = append(temp1, items[i+1:]...)
		doCombine1(temp1, temp, k, res)
	}
}

