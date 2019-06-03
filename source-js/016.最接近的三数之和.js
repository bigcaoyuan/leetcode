/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var threeSumClosest = function (nums, target) {
	let sortedNums = nums.sort((a, b) => a - b), finded = new Set, closest = Number.MAX_VALUE, ret, temp

	for (let k = 0; k < sortedNums.length - 2; k++) {
		if (finded.has(sortedNums[k]))
			continue

		let i = k + 1
			, j = sortedNums.length - 1

		while (i < j) {
			if (sortedNums[i] === sortedNums[i - 1] && sortedNums[j] === sortedNums[j + 1]) {
				j--
				continue
			}

			temp = sortedNums[i] + sortedNums[k] + sortedNums[j]

			if (temp === target) {
				return temp
			} else if (temp > target) {
				if (temp - target < closest) {
					closest = temp - target
					ret = temp
				}
				j--
			} else {
				if (target - temp < closest) {
					closest = target - temp
					ret = temp
				}
				i++
			}

		}

		finded.add(sortedNums[k])
	}

	return ret
};