import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
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
			sum := val + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{val, nums[l], nums[r]})

				for l < r {
					l++
					if nums[l] != nums[l-1] {
						break
					}
				}
				for l < r {
					r--
					if nums[r] != nums[r+1] {
						break
					}
				}

			} else if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			}
		}
	}
	return res
}

