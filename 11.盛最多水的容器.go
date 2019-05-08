/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
func maxArea(height []int) int {
	maxArea := 0
	l := 0
	r := len(height) - 1
	lastShort := 0
	tempShort := 0
	for {
		if l == r {
			break
		}
		if height[l] < height[r] {
			tempShort = height[l]
		} else {
			tempShort = height[r]
		}
		if lastShort < tempShort {
			tempArea := tempShort * (r - l)
			if tempArea > maxArea {
				maxArea = tempArea
			}
			lastShort = tempShort
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}

	}
	return maxArea
}

