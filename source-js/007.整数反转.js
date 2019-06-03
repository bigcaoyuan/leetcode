/**
 * @param {number} x
 * @return {number}
 */
var reverse = function (x) {
	let arr = [],
		i,
		j = x > 0 ? x : -x

	while (j > 0) {
		i = j % 10
		arr.push(i)

		j = (j - i) / 10
	}

	let ret = 0

	for (let k = 0; k < arr.length; k++) {
		ret += arr[k] * Math.pow(10, arr.length - 1 - k)
	}

	ret = x > 0 ? ret : -ret

	if (ret < -Math.pow(2, 31) || ret > Math.pow(2, 31) - 1) {
		ret = 0
	}

	return ret
};