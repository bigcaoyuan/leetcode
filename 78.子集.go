import (
	"math"
)

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */
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
			if i&cnt >= 1 {
				sub = append(sub, nums[j])
			}
		}
		res = append(res, sub)
	}
	return res
}

