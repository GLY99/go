package main

func canPartition(nums []int) bool {
	length := len(nums)
	if length < 2 {
		return false
	}
	sum := 0
	maxNum := nums[0]
	for _, num := range nums {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	if maxNum > target {
		return false
	}
	// dp[i][j] 0 - i是否存在一些数字和等于j
	dp := make([][]bool, length)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	for i := 0; i < length; i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < length; i++ {
		v := nums[i]
		for j := 1; j < target+1; j++ {
			if v <= j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[length-1][target]
}
