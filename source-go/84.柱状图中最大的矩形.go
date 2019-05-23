/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */
func largestRectangleArea(heights []int) int {
	hi := make([]int, len(heights))
	sum := make([]int, len(heights))
	hi[0] = heights[0]
	sum[0] = heights[0]
	for i := 1; i < len(heights); i++ {
		if heights[i] > hi[i]-1 {

		}
	}
}

