/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
	let map = new Map

	for (let i = 0; i < nums.length; i++) {
		const needNum = target - nums[i]
		const needNumIndex = map.get(needNum)

		// 这里用map.has(needNum)而不直接用needNumIndex检测是否存在，因为needNumIndex有时候是0
		if (map.has(needNum) && needNumIndex !== i) {
			return [needNumIndex, i]
		}

		map.set(nums[i], i)
	}
};