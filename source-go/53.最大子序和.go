/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */
func maxSubArray(nums []int) int {
	maxhere, maxsum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if maxhere < 0 {
			maxhere = nums[i]
			if nums[i] > maxsum {
				maxsum = nums[i]
			}
		} else {
			maxhere = maxhere + nums[i]
			if maxhere > maxsum {
				maxsum = maxhere
			}
		}
	}
	return maxsum
}

