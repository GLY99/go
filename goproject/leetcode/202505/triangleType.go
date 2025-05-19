package main

// https://leetcode.cn/problems/type-of-triangle/?envType=daily-question&envId=2025-05-19

import "sort"

func triangleType(nums []int) string {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	if nums[0]+nums[1] <= nums[2] {
		return "none"
	} else if nums[0] == nums[2] {
		return "equilateral"
	} else if nums[1] == nums[0] || nums[1] == nums[2] {
		return "isosceles"
	} else {
		return "scalene"
	}
}
