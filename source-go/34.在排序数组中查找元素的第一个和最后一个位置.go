/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	l := 0
	r := len(nums) - 1
	for {
		if l > r {
			break
		}
		mid := (l + r) / 2
		if nums[mid] == target {
			res = []int{mid, mid}
			break
		}
		if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if res[0] == -1 {
		return res
	}

	l1 := l
	r1 := (l+r)/2 - 1
	for {
		if l1 > r1 {
			break
		}
		mid := (l1 + r1) / 2
		if nums[mid] == target {
			res[0] = mid
			r1 = r1 - 1
			continue
		}
		l1 = mid + 1
	}

	l2 := (l+r)/2 + 1
	r2 := r
	for {
		if l2 > r2 {
			break
		}
		mid := (l2 + r2) / 2
		if nums[mid] == target {
			res[1] = mid
			l2 = l2 + 1
			continue
		}
		r2 = mid - 1
	}

	return res
}

