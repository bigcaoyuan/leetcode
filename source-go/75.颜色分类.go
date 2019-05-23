import "fmt"

/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// 将0置于队列头， 将2置于队列位即可
func sortColors(nums []int) {
	endRed, startBlue := 0, len(nums)-1
	for i := 0; i < len(nums); {
		fmt.Println(i, nums)
		if nums[i] == 0 {
			if i != endRed {
				nums[i], nums[endRed] = nums[endRed], nums[i]
				endRed++
			} else {
				endRed++
				i++
			}
		} else if nums[i] == 2 {
			if i > startBlue {
				break
			} else {
				nums[i], nums[startBlue] = nums[startBlue], nums[i]
				startBlue--
			}
		} else {
			i++
		}
	}
}

// 太过于麻烦
func sortColors1(nums []int) {
	startRed, endRed, startWhite, endWhite, startBlue, endBlue := -1, -1, -1, -1, -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			// 之前没有红色
			if startRed == -1 {
				// 之前没有别的颜色的球
				if startWhite == -1 && startBlue == -1 {
					startRed, endRed = i, i
				} else if startWhite == -1 {
					startRed, endRed = 0, 0
					startBlue, endBlue = startBlue+1, endBlue+1
					nums[0], nums[i] = nums[i], nums[0]
				} else if startBlue == -1 {
					startRed, endRed = 0, 0
					startWhite, endWhite = startWhite+1, endWhite+1
					nums[0], nums[i] = nums[i], nums[0]
				} else {
					startRed, endRed = 0, 0
					startWhite, endWhite = startWhite+1, endWhite+1
					startBlue, endBlue = startBlue+1, endBlue+1
					nums[0] = 0
					nums[endWhite] = 1
					nums[endBlue] = 2
				}
			} else {
				// 之前没有别的颜色的球
				if startWhite == -1 && startBlue == -1 {
					endRed = i
				} else if startWhite == -1 {
					endRed = endRed + 1
					startBlue, endBlue = startBlue+1, endBlue+1
					nums[endRed], nums[i] = nums[i], nums[endRed]
				} else if startBlue == -1 {
					endRed = endRed + 1
					startWhite, endWhite = startWhite+1, endWhite+1
					nums[endRed], nums[i] = nums[i], nums[endRed]
				} else {
					endRed = endRed + 1
					startWhite, endWhite = startWhite+1, endWhite+1
					startBlue, endBlue = startBlue+1, endBlue+1
					nums[endRed] = 0
					nums[endWhite] = 1
					nums[endBlue] = 2
				}
			}
		} else if nums[i] == 1 {
			// 之前没有出现过白色
			if endWhite == -1 {
				// 也没出现过蓝色
				if startBlue == -1 {
					startWhite, endWhite = i, i
				} else {
					// 之前出现过蓝色，将第一个蓝色和当前位置的白色对调即可
					startWhite, endWhite = startBlue, startBlue
					startBlue, endBlue = startBlue+1, endBlue+1
					nums[startWhite], nums[endBlue] = nums[endBlue], nums[startWhite]
				}
			} else {
				nums[endWhite+1], nums[i] = nums[i], nums[endWhite+1]
				endWhite = endWhite + 1
				if startBlue != -1 {
					startBlue, endBlue = startBlue+1, endBlue+1
				}
			}
		} else if nums[i] == 2 {
			// 出现蓝色其实不需要调动位置，蓝色本身就在最后面
			endBlue = i
			// 之前没有出现过蓝色
			if startBlue == -1 {
				startBlue = i
			}
		}
	}
}

