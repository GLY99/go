package main

import "sort"

// https://leetcode.cn/problems/divide-an-array-into-subarrays-with-minimum-cost-i/?envType=problem-list-v2&envId=sorting

func minimumCost(nums []int) int {
	ans := nums[0]
	nums = nums[1:]
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	ans += nums[0]
	ans += nums[1]
	return ans
}
