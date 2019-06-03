/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function (height) {
	let i = 0, j = height.length - 1, max = 0

	while (i < j) {
		if (height[i] < height[j]) {
			max = Math.max(max, height[i] * (j - i))
			i++
		} else {
			max = Math.max(max, height[j] * (j - i))
			j--
		}
	}

	return max
};