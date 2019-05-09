/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */
func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	lastIndex := 0
	lastValue := nums[0]
	isDump := false
	for i := 1; i < len(nums); i++ {
		// 发现重复元素
		if nums[i] == lastValue {
			// 元素还未重复过
			if !isDump {
				isDump = true
				lastIndex++
				nums[lastIndex] = nums[i]
			}
		} else {
			lastValue = nums[i]
			lastIndex++
			nums[lastIndex] = nums[i]
			isDump = false
		}
	}
	return lastIndex + 1
}

