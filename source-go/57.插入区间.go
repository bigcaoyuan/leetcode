/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}
	if newInterval[0] > intervals[len(intervals)-1][1] {
		return append(intervals, newInterval)
	}
	first := 0
	for i := 0; i < len(intervals); i++ {
		temp := intervals[i]
		if newInterval[0] > temp[1] {
			continue
		} else {
			// 插入
			if newInterval[1] < temp[0] {
				temp1 := append([][]int{newInterval}, intervals[i:]...)
				// temp1 := append(intervals[:i], newInterval)
				intervals = append(intervals[:i], temp1...)
				return intervals
			} else {
				if newInterval[0] < temp[0] {
					intervals[i][0] = newInterval[0]
				}
			}
			first = i
			break
		}
	}
	second := first
	for i := first; i < len(intervals); i++ {
		temp := intervals[i]
		if newInterval[1] > temp[1] {
			if i == len(intervals)-1 {
				intervals[i][1] = newInterval[1]
				second = len(intervals) - 1
			}
			continue
		}
		if newInterval[1] <= temp[1] && newInterval[1] >= temp[0] {
			second = i
			break
		}
		if newInterval[1] < temp[0] {
			second = i - 1
			if second < first {
				second = first
			}
			if newInterval[1] > intervals[second][1] {
				intervals[second][1] = newInterval[1]
			}
			break
		}
	}
	if first == second {
		return intervals
	}
	intervals[first][1] = intervals[second][1]
	intervals = append(intervals[:first+1], intervals[second+1:]...)
	return intervals
}

