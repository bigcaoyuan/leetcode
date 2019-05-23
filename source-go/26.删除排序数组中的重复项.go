/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	cnt := 1
	lastVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != lastVal {
			lastVal = nums[i]
			nums[cnt] = nums[i]
			cnt++
		}
	}
	return cnt
}

