/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[][]}
 */
var fourSum = function (nums, target) {
	let sortedNums = nums.sort((a, b) => a - b)
		, ret = []

	let i = 0, j

	for (let i = 0; i < sortedNums.length - 3; i++) {
		if (i > 0 && sortedNums[i] === sortedNums[i - 1])
			continue

		for (let j = i + 1; j < sortedNums.length - 2; j++) {
			if (j > i + 1 && sortedNums[j] === sortedNums[j - 1])
				continue

			let k = j + 1
				, l = sortedNums.length - 1

			while (k < l) {
				if (sortedNums[k] === sortedNums[k - 1] && sortedNums[l] === sortedNums[l + 1]) {
					l--
					continue
				}
				if (sortedNums[i] + sortedNums[j] + sortedNums[k] + sortedNums[l] === target) {
					ret.push([sortedNums[i], sortedNums[j], sortedNums[k], sortedNums[l]])
					l--;
					k++;
				} else if (sortedNums[i] + sortedNums[j] + sortedNums[k] + sortedNums[l] > target) {
					l--
				} else {
					k++
				}
			}
		}
	}

	return ret
};