package main

// https://leetcode.cn/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	length := len(nums)
	ans := nums[0]
	dp := make([]int, length) // dp[i]表示以第i个元素结尾的最大值
	dp[0] = nums[0]
	for i := 1; i < length; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(dp[i], ans)
	}
	return ans
}

func maxSubArray1(nums []int) int {
	length := len(nums)
	ans := nums[0]
	curMax := ans // 当前状态之和前一个状态有关，只维护前一个状态
	for i := 1; i < length; i++ {
		curMax = max(curMax+nums[i], nums[i])
		ans = max(curMax, ans)
	}
	return ans
}
