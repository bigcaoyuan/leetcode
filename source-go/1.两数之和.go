/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if val, ok := mp[diff]; ok && val != i {
			return []int{i, val}
		}
	}
	return []int{}
}

