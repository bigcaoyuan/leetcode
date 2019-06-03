/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function (nums1, nums2) {
	let i1 = 0,
		i2 = 0,
		result = []

	const len1 = nums1.length,
		len2 = nums2.length

	while (i1 < len1 && i2 < len2) {
		if (nums1[i1] <= nums2[i2]) {
			result.push(nums1[i1++])
		} else {
			result.push(nums2[i2++])
		}
	}

	while (i1 < len1) {
		result.push(nums1[i1++])
	}
	while (i2 < len2) {
		result.push(nums2[i2++])
	}

	const lenr = result.length

	if (lenr % 2 === 0) {
		return (result[(lenr / 2) - 1] + result[lenr / 2]) / 2
	} else {
		return result[Math.floor(lenr / 2)]
	}
};