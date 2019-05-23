import "fmt"

/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 */
func getPermutation(n int, k int) string {
	ns := make([]int, n)
	for i := 0; i < len(ns); i++ {
		ns[i] = i + 1
	}
	res := 0

	for {
		if len(ns) == 1 {
			res = res*10 + ns[0]
			break
		}
		temp := getN(len(ns) - 1)
		s1 := k / temp
		s2 := k % temp
		if s2 == 0 {
			s1 = s1 - 1
			if s1 < 0 {
				s1 = 0
			}
		}
		res = res*10 + ns[s1]
		ns = append(ns[:s1], ns[s1+1:]...)
		k = k - temp*s1
		if k == 0 {
			for i := 0; i < len(ns); i++ {
				res = res*10 + ns[i]
			}
			break
		}
	}
	return fmt.Sprintf("%d", res)
}

func getN(k int) int {
	res := 1
	for i := 2; i <= k; i++ {
		res *= i
	}
	return res
}

