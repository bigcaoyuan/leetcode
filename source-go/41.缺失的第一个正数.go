/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */
func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	for i := 0; i < len(nums); i++ {
		j := i
		for {
			if nums[j] == j+1 || nums[j] < 1 || nums[j] > len(nums) {
				break
			}
			if nums[nums[j]-1] == nums[j] {
				nums[j] = 0
				break
			}
			nums[j], nums[nums[j]-1] = nums[nums[j]-1], nums[j]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

