package main

func getCommon(nums1 []int, nums2 []int) int {
	idx1 := 0
	idx2 := 0
	for idx1 < len(nums1) && idx2 < len(nums2) {
		for idx1 < len(nums1) && nums1[idx1] < nums2[idx2] {
			idx1++
		}
		if idx1 >= len(nums1) {
			break
		}
		if nums1[idx1] == nums2[idx2] {
			return nums1[idx1]
		}
		for idx2 < len(nums2) && nums2[idx2] < nums1[idx1] {
			idx2++
		}
		if idx2 >= len(nums2) {
			break
		}
		if nums1[idx1] == nums2[idx2] {
			return nums1[idx1]
		}
	}
	return -1
}
