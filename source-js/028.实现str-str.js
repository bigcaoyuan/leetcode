function getNext(needle) {
	let next = [-1]
		, i = 0
		, j = -1

	while (i < needle.length - 1) {
		if (j === -1 || needle[i] === needle[j]) {
			next[++i] = ++j
		} else {
			j = next[j]
		}
	}

	return next
}

/**
* @param {string} haystack
* @param {string} needle
* @return {number}
*/
var strStr = function (haystack, needle) {
	if (needle.length === 0)
		return 0

	let i = 0
		, j = 0
		, next = getNext(needle)

	while (i < haystack.length && j < needle.length) {
		if (j === -1 || haystack[i] === needle[j]) {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if (j === needle.length) {
		return i - j
	}

	return -1
};