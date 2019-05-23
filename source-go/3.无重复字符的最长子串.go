/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
func lengthOfLongestSubstring(s string) int {

	// copy_s := []rune(s)
	// curr_list := []rune{}
	// max_len := 0
	// for _, v := range copy_s {
	//     for sub_i, sub_v := range curr_list {
	//         if v == sub_v {
	//             curr_list = curr_list[sub_i+1: ]
	//             break
	//         }
	//     }
	//     curr_list = append(curr_list, v)
	//     if len(curr_list) > max_len {
	//         max_len = len(curr_list)
	//     }
	// }
	// return max_len

	uniqMp := make(map[byte]int)
	maxsum := 0
	for i := 0; i < len(s); i++ {
		if index, ok := uniqMp[s[i]]; !ok {
			uniqMp[s[i]] = i
			if len(uniqMp) > maxsum {
				maxsum = len(uniqMp)
			}
		} else {
			uniqMp[s[i]] = i
			for key, val := range uniqMp {
				if val <= index {
					delete(uniqMp, key)
				}
			}
		}
	}
	return maxsum
}

