/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function (nums, val) {
	if (nums.length === 0) return 0

	let i = 0
		, j = 0

	for (; j < nums.length; j++) {
		if (nums[j] !== val) {
			nums[i++] = nums[j]
		}
	}

	return i
};