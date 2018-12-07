func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxCnt := 1
	for i := 0; i < len(s); i++ {
		keyMap := make(map[byte]int)
		for j := i; j < len(s); j++ {
			if _, ok := keyMap[s[j]]; !ok {
				keyMap[s[j]] = j
				cnt := j + 1 - i
				if cnt > maxCnt {
					maxCnt = cnt
				}
				continue
			}
			break
		}
	}
	return maxCnt
}