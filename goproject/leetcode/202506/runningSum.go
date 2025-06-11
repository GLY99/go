package main

// https://leetcode.cn/problems/running-sum-of-1d-array/submissions/636112560/?envType=problem-list-v2&envId=array

func runningSum(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] + nums[i]
	}
	return ans
}
