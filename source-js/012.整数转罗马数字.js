/**
 * @param {number} num
 * @return {string}
 */
var intToRoman = function (num) {
	const values = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]
	const strs = ['M', 'CM', 'D', 'CD', 'C', 'XC', 'L', 'XL', 'X', 'IX', 'V', 'IV', 'I']

	let ret = ''
	for (let i = 0; i < values.length; i++) {
		while (num >= values[i]) {
			ret += strs[i]
			num -= values[i]

			if (num === 0) return ret
		}
	}

	return ret
};