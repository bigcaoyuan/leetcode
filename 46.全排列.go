/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */
func permute(nums []int) [][]int {
	res := [][]int{}
	if len(nums) <= 1 {
		return append(res, nums)
	}
	p(&res, []int{}, nums)
	return res
}

func p(res *[][]int, temp, rest []int) {
	if len(rest) == 0 {
		tempp := make([]int, len(temp))
		copy(tempp, temp)
		*res = append(*res, tempp)
		return
	}
	for i := 0; i < len(rest); i++ {
		temp1 := append(temp, rest[i])
		r := append([]int{}, rest[:i]...)
		r = append(r, rest[i+1:]...)
		p(res, temp1, r)
	}
}

