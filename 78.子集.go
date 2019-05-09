import "math"

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// 质数法 受评论启发
func subsets(nums []int) [][]int {
	res := [][]int{}
	n := math.Pow(2, float64(len(nums)))
	nint := int(n)
	for i := 0; i < nint; i++ {
		sub := []int{}
		cnt := 1
		for j := 0; j < len(nums); j++ {
			if j == 0 {
				cnt = 1
			} else {
				cnt = cnt << 1
			}
			if i&cnt > 0 {
				sub = append(sub, nums[j])
			}
		}
		res = append(res, sub)
	}
	return res
}

// 回溯法
func subsets1(nums []int) [][]int {
	res := [][]int{[]int{}}
	if len(nums) == 0 {
		return res
	}
	for i := 0; i < len(nums); i++ {
		doSubsets1(nums, []int{nums[i]}, i+1, &res)
	}
	return res
}
func doSubsets1(items []int, local []int, index int, res *[][]int) {
	*res = append(*res, local)
	for i := index; i < len(items); i++ {
		temp := append([]int{}, local...)
		temp = append(temp, items[i])
		doSubsets1(items, temp, i+1, res)
	}
}

