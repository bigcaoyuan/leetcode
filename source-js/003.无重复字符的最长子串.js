/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function (s) {
	let max = 0, charSet = new Set

	for (let i = 0; i < s.length; i++) {
		if (charSet.has(s[i])) {
			max = Math.max(max, charSet.size)
			// 删掉最左侧字符，继续遍历当前项
			charSet.delete(s[i - charSet.size])
			i--
		} else {
			charSet.add(s[i])
		}

	}
	return Math.max(max, charSet.size)
}