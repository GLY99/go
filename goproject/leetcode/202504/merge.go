package main

// https://leetcode.cn/problems/merge-sorted-array/submissions/621036041/?envType=problem-list-v2&envId=sorting
// 原地合并两个有序数组，思路是从后合并

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for i >= 0 {
		nums1[k] = nums1[i]
		k--
		i--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}
