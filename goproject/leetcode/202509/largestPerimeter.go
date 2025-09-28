package main

import "sort"

// https://leetcode.cn/problems/largest-perimeter-triangle/?envType=daily-question&envId=2025-09-28

func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	n := len(nums)
	for i := n - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i-2] + nums[i-1] + nums[i]
		}
	}
	return 0
}
