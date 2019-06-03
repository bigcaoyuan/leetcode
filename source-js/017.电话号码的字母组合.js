/**
 * @param {string} digits
 * @return {string[]}
 */
var map = {
	2: 'abc',
	3: 'def',
	4: 'ghi',
	5: 'jkl',
	6: 'mno',
	7: 'pqrs',
	8: 'tuv',
	9: 'wxyz',
}
var letterCombinations = function (digits) {
	let ret = []

	if (digits.length > 1) {
		let left = digits.slice(0, digits.length - 1)
		let right = digits[digits.length - 1]

		const subDigits = letterCombinations(left)

		for (let i = 0; i < subDigits.length; i++) {
			for (let j = 0; j < map[right].length; j++) {
				ret.push(subDigits[i] + map[right][j])
			}
		}
	} else if (digits.length === 1) {
		for (let j = 0; j < map[digits].length; j++) {
			ret.push(map[digits][j])
		}
	}

	return ret
};