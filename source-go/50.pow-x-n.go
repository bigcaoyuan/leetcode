/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */
func myPow(x float64, n int) float64 {
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	res := pow(x, n)
	if neg {
		res = 1 / res
	}
	return res
}

func pow(x float64, n int) float64 {
	if n/2 == 0 {
		if n == 1 {
			return x
		} else {
			return 1
		}
	}

	if n%2 == 0 {
		res := pow(x, n/2)
		return res * res
	} else {
		res := pow(x, n/2)
		return res * res * x
	}
}
