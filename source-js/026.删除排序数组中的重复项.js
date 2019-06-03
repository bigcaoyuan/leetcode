/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function (nums) {
	if (nums.length < 2)
		return nums.length

	let i = 0
		, j = 1

	for (; j < nums.length; j++) {
		if (nums[i] !== nums[j])
			nums[++i] = nums[j]
	}

	return i + 1
};