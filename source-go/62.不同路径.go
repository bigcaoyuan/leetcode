/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 * 这个问题看似是阶乘(全组合)问题  但实际不是，如果阶乘完了可以去重也可以，但很复杂
 */
func uniquePaths(m int, n int) int {
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
		cache[i][0] = 1
	}
	for i := 0; i < n; i++ {
		cache[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cache[i][j] = cache[i-1][j] + cache[i][j-1]
		}
	}
	return cache[m-1][n-1]
}

// func uniquePaths(m int, n int) int {
// 	if m == 1 && n == 1 {
// 		return 1
// 	}
// 	cache := make(map[int]map[int]int)
// 	return doUniquePaths(m-1, n, cache) + doUniquePaths(m, n-1, cache)
// }

// func doUniquePaths(m, n int, cache map[int]map[int]int) int {
// 	if m == 0 || n == 0 {
// 		return 0
// 	}
// 	if n < m {
// 		m, n = n, m
// 	}
// 	if _, ok := cache[m]; !ok {
// 		cache[m] = make(map[int]int)
// 	}
// 	if _, ok := cache[m][n]; ok {
// 		return cache[m][n]
// 	}
// 	if m > 1 && n > 1 {
// 		res := doUniquePaths(m-1, n, cache) + doUniquePaths(m, n-1, cache)
// 		cache[m][n] = res
// 		return res
// 	}
// 	return 1
// }

