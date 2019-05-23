import "sort"

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		left := nums[i]
		if i != 0 && left == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			right := nums[j]
			if j != i+1 && right == nums[j-1] {
				continue
			}

			l := j + 1
			r := len(nums) - 1
			for {
				if l >= r {
					break
				}
				sum := left + right + nums[l] + nums[r]
				if sum == target {
					res = append(res, []int{left, right, nums[l], nums[r]})
					l++
					for nums[l] == nums[l-1] && l < r {
						l++
					}
					r--
					for nums[r] == nums[r+1] && l < r {
						r--
					}
				} else if sum < target {
					l++
				} else if sum > target {
					r--
				}
			}
		}
	}
	return res
}

