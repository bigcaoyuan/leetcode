/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	for cur := m + n - 1; cur >= 0; cur-- {
		if i < 0 {
			nums1[cur] = nums2[j]
			j--
			continue
		} else if j < 0 {
			return
		}
		if nums1[i] > nums2[j] {
			nums1[cur] = nums1[i]
			i--
		} else {
			nums1[cur] = nums2[j]
			j--
		}
	}
}

