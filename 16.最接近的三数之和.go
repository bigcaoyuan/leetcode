import "sort"

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minNum := nums[0] + nums[1] + nums[2]
	minAbsolute := minNum - target
	if minAbsolute < 0 {
		minAbsolute = -minAbsolute
	}
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		val := nums[i]
		l := i + 1
		r := len(nums) - 1
		for {
			if l >= r {
				break
			}
			sum := val + nums[l] + nums[r] - target
			if sum == 0 {
				return target
			} else if sum > 0 {
				if sum < minAbsolute {
					minAbsolute = sum
					minNum = sum + target
				}
				r--
			} else if sum < 0 {
				l++
				if -sum < minAbsolute {
					minAbsolute = -sum
					minNum = sum + target
				}
			}
		}
	}
	return minNum
}

