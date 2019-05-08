/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 * 动态规划题
 * 另一种技巧，数组中没有0则一定可以到达最后
 */
func canJump(nums []int) bool {
	// return canJump1(nums)
	return canJump2(nums)
}

func canJump1(nums []int) bool {
	maxDist := nums[0]
	for i := 1; i < len(nums) && i <= maxDist; i++ {
		if nums[i]+i > maxDist {
			maxDist = nums[i] + i
		}
		if maxDist >= len(nums)-1 {
			return true
		}
	}
	if maxDist >= len(nums)-1 {
		return true
	}
	return false
}

func canJump2(nums []int) bool {
	lastReach := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] != 0 && nums[i]+i >= lastReach {
			lastReach = i
		}
	}

	if lastReach == 0 {
		return true
	}
	return false
}

