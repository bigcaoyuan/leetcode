/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	for i := 0; i < len(intervals); i++ {
		min := i
		for j := i; j < len(intervals); j++ {
			if intervals[j][0] < intervals[min][0] {
				min = j
			}
		}
		intervals[i], intervals[min] = intervals[min], intervals[i]
	}

	res := make([][]int, 0)
	start := intervals[0][0]
	end := intervals[0][1]
	index := 1
	for {
		if index == len(intervals) {
			res = append(res, []int{start, end})
			break
		}

		if end >= intervals[index][0] {
			// end = intervals[index][1]
			end = max(end, intervals[index][1])
			start = min(start, intervals[index][0])
			index++
			continue
		}

		res = append(res, []int{start, end})
		start = intervals[index][0]
		end = intervals[index][1]
		index++
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
