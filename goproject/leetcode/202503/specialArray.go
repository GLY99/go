package main

import "sort"

// https://leetcode.cn/problems/special-array-with-x-elements-greater-than-or-equal-x/?envType=problem-list-v2&envId=sorting

func specialArray(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	length := len(nums)
	for i := 1; i <= length; i++ {
		if nums[i-1] >= i && (i == length || nums[i] < i) {
			return i
		}
	}
	return -1
}
