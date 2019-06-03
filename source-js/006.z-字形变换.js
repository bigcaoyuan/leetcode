/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
var convert = function (s, numRows) {
	let zArr = [],
		row = 0,
		// 行号加减方向
		reverse = true

	for (let j = 0; j < s.length; j++) {
		if (zArr[row]) {
			zArr[row].push(s[j])
		} else {
			zArr[row] = [s[j]]
		}

		if (row === numRows - 1 || row === 0) {
			reverse = !reverse
		}

		reverse ? row-- : row++
	}

	let result = ''

	for (let k = 0; k < zArr.length; k++) {
		for (let l = 0; l < zArr[k].length; l++) {
			result += zArr[k][l]
		}
	}

	return result;
};