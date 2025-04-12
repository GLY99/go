package main

import "sort"

// https://leetcode.cn/problems/WhsWhI/
// https://leetcode.cn/problems/longest-consecutive-sequence/description/

func longestConsecutive(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	startIdx := 0
	endIdx := 1
	repeatCount := 0
	ans := 0
	for endIdx < len(nums) {
		if nums[endIdx] != nums[endIdx-1]+1 && nums[endIdx] != nums[endIdx-1] {
			ans = max(ans, endIdx-startIdx-repeatCount)
			startIdx = endIdx
			repeatCount = 0
		} else if nums[endIdx] != nums[endIdx-1] {
			repeatCount++
		}
		endIdx++
	}
	return max(ans, endIdx-startIdx-repeatCount)
}
