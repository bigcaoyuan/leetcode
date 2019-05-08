/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */
func removeElement(nums []int, val int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[cnt] = nums[i]
			cnt++
		}
	}
	return cnt
}

