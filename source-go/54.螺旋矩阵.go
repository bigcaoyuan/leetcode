/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	res := make([]int, len(matrix)*len(matrix[0]))
	i, j, index, cnt, right, row := 0, 0, 0, 0, true, true

	for {

		if cnt == len(res) {
			break
		}
		if row && right {
			if j == len(matrix[0])-index {
				row = false
				j--
				i++
				continue
			}
			res[cnt] = matrix[i][j]
			cnt++
			j++
		}
		if !row && right {

			if i == len(matrix)-index {
				row = true
				right = false
				i--
				j--
				continue
			}
			res[cnt] = matrix[i][j]
			cnt++
			i++
		}

		if row && !right {

			if j == index-1 {
				row = false
				j++
				i--
				continue
			}
			res[cnt] = matrix[i][j]
			cnt++
			j--

		}

		if !row && !right {
			if i == index {
				row = true
				right = true
				i++
				j++
				index++
				continue
			}
			res[cnt] = matrix[i][j]
			cnt++
			i--
		}
	}
	return res
}

// iIndex, cnt, rowCnt, colCnt := 0, 0, len(matrix), len(matrix[0])
// for {
// 	if iIndex > rowCnt/2 {
// 		break
// 	}
// 	temp := colCnt - iIndex*2
// 	for i := 0; i < temp; i++ {
// 		res[cnt] = matrix[iIndex][iIndex+i]
// 		cnt++
// 	}

// 	temp = rowCnt - iIndex*2
// 	for i := 1; i < temp && colCnt-1-iIndex != iIndex; i++ {
// 		res[cnt] = matrix[iIndex+i][colCnt-1-iIndex]
// 		cnt++
// 	}

// 	temp = colCnt - iIndex*2
// 	for i := 1; i < temp && rowCnt-1-iIndex != iIndex; i++ {
// 		res[cnt] = matrix[rowCnt-1-iIndex][colCnt-1-iIndex-i]
// 		cnt++
// 	}

// 	temp = rowCnt - iIndex*2
// 	for i := 1; i < temp-1; i++ {
// 		res[cnt] = matrix[rowCnt-1-iIndex-i][iIndex]
// 		cnt++
// 	}
// 	iIndex++

// }
// return res


