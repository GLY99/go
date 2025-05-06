package main

// https://leetcode.cn/problems/build-array-from-permutation/?envType=daily-question&envId=2025-05-06

func buildArray(nums []int) []int {
	length := len(nums)
	ans := make([]int, length)
	for i := 0; i < length; i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}
