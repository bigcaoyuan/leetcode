/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */
func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	mp := make(map[int]int)
	mp[0] = 1
	mp[1] = 1
	mp[2] = 2
	mp[3] = 5
	for i := 4; i <= n; i++ {
		// j 表示左子树元素个数
		// i-1-j表示右子树元素个数
		for j := 0; j < i; j++ {
			left := mp[j]
			right := mp[i-1-j]
			mp[i] += left * right
		}
	}
	return mp[n]
}


