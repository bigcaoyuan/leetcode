/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// 栈思想
func trap(height []int) int {
	stack := make([]int, 0)
	res := 0
	for i := 0; i < len(height); i++ {
		for {
			if len(stack) == 0 || height[i] < height[stack[len(stack)-1]] {
				break
			}
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			lastIndex := stack[len(stack)-1]

			minHeight := height[i]
			if height[lastIndex] < minHeight {
				minHeight = height[lastIndex]
			}
			res += (minHeight - height[tmp]) * (i - lastIndex - 1)
		}
		stack = append(stack, i)
	}
	return res
}

// 双向扫描空白区域 最后求反值即可
func trap1(height []int) int {
	max := 0
	for _, length := range height {
		if length > max {
			max = length
		}
	}
	start := make([]int, len(height))
	end := make([]int, len(height))
	temp := max
	for i := 0; i < len(height); i++ {
		if max-height[i] < temp {
			temp = max - height[i]
		}
		start[i] = temp
		if temp == 0 {
			break
		}
	}
	temp = max
	for i := len(height) - 1; i >= 0; i-- {
		if max-height[i] < temp {
			temp = max - height[i]
		}
		end[i] = temp
		if temp == 0 {
			break
		}
	}
	sum := 0
	for i := 0; i < len(height); i++ {
		if start[i] > end[i] {
			sum += start[i]
		} else {
			sum += end[i]
		}
		sum += height[i]
	}
	sum = len(height)*max - sum
	return sum
}




