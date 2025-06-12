package main

// https://leetcode.cn/problems/maximum-difference-between-adjacent-elements-in-a-circular-array/?envType=daily-question&envId=2025-06-12

func maxAdjacentDistance(nums []int) int {
	length := len(nums)
	ans := abs(nums[0] - nums[length-1])
	for i := 1; i < len(nums); i++ {
		ans = max(ans, abs(nums[i]-nums[i-1]))
	}
	return ans
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
