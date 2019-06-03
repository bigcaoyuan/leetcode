function isInterger(c) {
	return c === '0' || c === '1' || c === '2' || c === '3' || c === '4' || c === '5' || c === '6' || c === '7' || c === '8' || c === '9'
}
/**
 * @param {string} str
 * @return {number}
 */
var myAtoi = function (str) {
	let i = 0, retArr = [], ret = 0, sign = false
	const INT_MAX = Math.pow(2, 31) - 1,
		INT_MIN = -Math.pow(2, 31)

	while (str[i] === ' ' && i < str.length) {
		i++
	}

	if (i === str.length) return 0

	if (str[i] === '+' || str[i] === '-' || isInterger(str[i])) {
		if (!isInterger(str[i])) {
			sign = str[i] === '+' ? false : true
			i++
		}
	} else {
		return 0
	}

	while (isInterger(str[i]) && i < str.length) {
		retArr.push(str[i++])
	}

	if (retArr.length === 0) return 0

	for (let j = 0; j < retArr.length; j++) {
		ret += retArr[j] * Math.pow(10, retArr.length - 1 - j)
	}

	ret = sign ? -ret : ret

	if (ret > INT_MAX) return INT_MAX
	if (ret < INT_MIN) return INT_MIN

	return ret
};