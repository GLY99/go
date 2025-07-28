package main

// https://leetcode.cn/problems/sum-of-all-subset-xor-totals/?envType=problem-list-v2&envId=array

func subsetXORSum(nums []int) int {
	return dfs(0, 0, nums)
}

func dfs(val, idx int, nums []int) int {
	if idx == len(nums) {
		return val
	}
	return dfs(val^nums[idx], idx+1, nums) + dfs(val, idx+1, nums)
}
