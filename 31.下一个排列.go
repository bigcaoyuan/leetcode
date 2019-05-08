/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	index := 0
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			index = i
			break
		}
	}
	if index != 0 {
		lastVal := nums[index-1]
		for j := index; j < len(nums)-1; j++ {
			if nums[j] > nums[index-1] && nums[index-1] >= nums[j+1] {
				nums[index-1], nums[j] = nums[j], nums[index-1]
				break
			}
		}
		// 未变化
		if lastVal == nums[index-1] {
			nums[index-1], nums[len(nums)-1] = nums[len(nums)-1], nums[index-1]
		}
	}

	// 倒叙了
	length := len(nums) - 1
	for i := 0; ; i++ {
		if i+index >= length-i {
			break
		}
		nums[i+index], nums[length-i] = nums[length-i], nums[i+index]
	}
	return
}  

