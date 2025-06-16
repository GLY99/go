package main

// https://leetcode.cn/problems/maximum-difference-between-increasing-elements/?envType=daily-question&envId=2025-06-16

func maximumDifference(nums []int) int {
	ans := -1
	for i, preMin := 1, nums[0]; i < len(nums); i++ {
		if preMin >= nums[i] {
			preMin = nums[i]
		} else {
			ans = max(ans, nums[i]-preMin)
		}
	}
	return ans
}
