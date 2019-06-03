/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function (nums) {
	let ret = []
		, sortedNums = nums.sort((a, b) => a - b)
		, finded = new Set

	// 以最左侧元素作为基准元素，寻找两个与最左侧元素相加为0的数
	for (let k = 0; k < sortedNums.length - 2; k++) {
		// 基准元素没必要重复找
		if (finded.has(sortedNums[k]))
			continue

		let i = k + 1
			, j = sortedNums.length - 1

		while (i < j) {
			// 两个数也没必要重复找
			if (sortedNums[i] === sortedNums[i - 1] && sortedNums[j] === sortedNums[j + 1]) {
				j--
				continue
			}

			if (sortedNums[i] + sortedNums[k] + sortedNums[j] === 0) {
				ret.push([sortedNums[i], sortedNums[k], sortedNums[j]])
				j--
				i++
			} else if (sortedNums[i] + sortedNums[k] + sortedNums[j] > 0) {
				j--
			} else {
				i++
			}
		}

		finded.add(sortedNums[k])
	}

	return ret
};