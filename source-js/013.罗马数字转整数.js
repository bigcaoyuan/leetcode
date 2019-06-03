var map = {
	'I': 1,
	'V': 5,

	'X': 10,
	'L': 50,

	'C': 100,
	'D': 500,

	'M': 1000,
}
/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function (s) {
	let ret = 0

	for (let i = s.length - 1; i >= 0; i--) {

		if (s[i] === 'I') {
			if (s[i + 1] === 'V' || s[i + 1] === 'X') {
				ret -= 1
			} else {
				ret += 1
			}
		} else if (s[i] === 'X') {
			if (s[i + 1] === 'L' || s[i + 1] === 'C') {
				ret -= 10
			} else {
				ret += 10
			}
		} else if (s[i] === 'C') {
			if (s[i + 1] === 'D' || s[i + 1] === 'M') {
				ret -= 100
			} else {
				ret += 100
			}
		} else {
			ret += map[s[i]]
		}
	}

	return ret
};