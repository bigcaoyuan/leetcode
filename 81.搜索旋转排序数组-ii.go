/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */
func search(nums []int, target int) bool {
	return doSearch(nums, target, 0, len(nums)-1)
}

func doSearch(nums []int, target int, start, end int) bool {
	if start > end {
		return false
	}
	mid := (start + end) / 2
	if target == nums[mid] {
		return true
	}
	if nums[start] < nums[end] {
		// 未旋转
		if nums[mid] < target {
			return doSearch(nums, target, mid+1, end)
		} else {
			return doSearch(nums, target, start, mid-1)
		}
	} else {
		// 不确定
		if doSearch(nums, target, start, mid-1) {
			return true
		}
		return doSearch(nums, target, mid+1, end)
	}
	return false
}

