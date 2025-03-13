package main

// https://leetcode.cn/problems/count-number-of-pairs-with-absolute-difference-k/?envType=problem-list-v2&envId=hash-table

func countKDifference(nums []int, k int) int {
	mapping := map[int]int{}
	ans := 0
	for j := 0; j < len(nums); j++ {
		ans += mapping[nums[j]-k] + mapping[nums[j]+k]
		mapping[nums[j]]++
	}
	return ans
}
