/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */
func jump(nums []int) int {
	cnt := 0
	res := 0
	i := 0
	for {
		res = nums[i] + i + 1
		if res >= len(nums) {
			if i == len(nums)-1 {
				return cnt
			}
			return cnt + 1
		}
		cnt++
		tmp := i
		for j := 1; j <= nums[i]; j++ {
			if i+j+nums[i+j] > tmp+nums[tmp] {
				tmp = i + j
			}
		}
		i = tmp
	}
	return cnt
}

