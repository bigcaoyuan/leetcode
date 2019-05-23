/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */
func searchInsert(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if right == left {
		if target <= nums[right] {
			return right
		} else {
			return right + 1
		}
	}
	if right < left {
		if right < 0 {
			return 0
		}
		if left > len(nums)-1 {
			return len(nums)
		}

		if target < nums[right] {
			return right
		} else if target > nums[left] {
			return left + 1
		} else {
			return left
		}
	}
	return 0
}

