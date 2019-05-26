import "sort"

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */
func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums, []int{}}
	}
	if len(nums) == 0 {
		return [][]int{nums}
	}
	sort.Ints(nums)
	res := [][]int{[]int{}, []int{nums[0]}}
	lens := []int{1, 2}
	for i := 1; i < len(nums); i++ {
		left := 0
		if nums[i] == nums[i-1] {
			left = lens[i-1]
		}
		cnt := len(res)
		for j := left; j < cnt; j++ {
			temp := make([]int, len(res[j])+1)
			copy(temp, res[j])
			temp[len(temp)-1] = nums[i]
			res = append(res, temp)
		}
		lens = append(lens, len(res))
	}
	return res
}
