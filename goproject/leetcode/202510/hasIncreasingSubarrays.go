package main

// https://leetcode.cn/problems/adjacent-increasing-subarrays-detection-i/?envType=daily-question&envId=2025-10-14

func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)
	if n < 2*k {
		return false
	}
	if k == 1 {
		return true
	}
	for i := 0; i <= n-2*k; i++ {
		flag1 := true
		for j := i; j < i+k-1; j++ {
			if nums[j] >= nums[j+1] {
				flag1 = false
				break
			}
		}
		if !flag1 {
			continue
		}

		for j := i + k; j < i+2*k-1; j++ {
			if nums[j] >= nums[j+1] {
				flag1 = false
				break
			}
		}
		if flag1 {
			return true
		}
	}
	return false
}
