function maxPlalindrome(s, i) {
	let j = 1, k = 1, ret = s[i]

	while (k < s.length) {
		if (s[i] !== s[i + k]) break

		ret = ret + s[i + k]
		k++
	}

	// bb，寻找abba
	if (ret.length % 2 === 0) {
		while (i - j >= 0 && i + k - 1 + j < s.length) {
			if (s[i - j] !== s[i + k - 1 + j]) break

			ret = s[i - j] + ret + s[i + k - 1 + j]
			j++
		}
	}
	// b，寻找aba
	else {
		while (i - j >= 0 && i + k - 1 + j < s.length) {
			if (s[i - j] !== s[i + k - 1 + j]) break

			ret = s[i - j] + ret + s[i + k - 1 + j]

			j++
		}
	}

	return ret
}
/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function (s) {
	let ret = ''

	for (let i = 0; i < s.length; i++) {

		let plalindromet = maxPlalindrome(s, i)

		ret = plalindromet.length > ret.length ? plalindromet : ret
	}

	return ret
};