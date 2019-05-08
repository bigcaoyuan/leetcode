/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l := 0
	r := len(nums) - 1

	for {
		if l > r {
			break
		}
		mid := (l + r) / 2
		if target == nums[mid] {
			return mid
		}
		// 整个都未旋转
		if nums[l] <= nums[mid] && nums[mid] <= nums[r] {
			if target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
			continue
		}

		// 右边是未旋转的
		if nums[mid] < nums[r] {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
				continue
			}
			r = mid - 1
			continue
		}

		// 左边是未旋转的
		// if nums[l] < nums[mid] {
		if target < nums[mid] && target >= nums[l] {
			r = mid - 1
			continue
		}
		l = mid + 1
		continue
		// }

	}
	return -1
}

