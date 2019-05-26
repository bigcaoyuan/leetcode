import "sort"

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	temp := []int{}
	com(candidates, target, 0, temp, 0, &res)
	return res
}

func com(candidates []int, target, index int, temp []int, sum int, res *[][]int) {
	if sum == target {
		temp1 := make([]int, len(temp))
		copy(temp1, temp)
		*res = append(*res, temp1)
		return
	}
	if sum > target {
		return
	}
	last := -1
	for i := index; i < len(candidates); i++ {
		if candidates[i] == last {
			continue
		}
		last = candidates[i]
		temp1 := append(temp, candidates[i])
		sum1 := sum + candidates[i]
		com(candidates, target, i+1, temp1, sum1, res)
	}
}

