/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */
func groupAnagrams(strs []string) [][]string {
	// 想到了用二进制的且  则 两个单词且完了与自身相同
	// 但且不能去重 如 etaa eta  或者 etaaa eteee
	// 可以采用质数的形式

	prims := []int{
		2, 3, 5, 7, 11, 13, 17, 19,
		23, 29, 53, 59, 83, 89,
		31, 37, 61, 67,
		41, 43, 47, 71, 73,
		101, 103, 107, 109,
		131, 133, 137, 139,
		161, 163, 167, 169,
		191, 193, 197, 199,
	}

	mp := make(map[int][]string)
	for i := 0; i < len(strs); i++ {
		sum := 1
		for j := 0; j < len(strs[i]); j++ {
			sum *= prims[strs[i][j]-'a']
		}
		if _, ok := mp[sum]; !ok {
			mp[sum] = make([]string, 0)
		}
		mp[sum] = append(mp[sum], strs[i])
	}

	res := make([][]string, len(mp))
	cnt := 0
	for _, val := range mp {
		res[cnt] = val
		cnt++
	}
	return res
}

