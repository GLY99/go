package main

// https://leetcode.cn/problems/longest-harmonious-subsequence/?envType=daily-question&envId=2025-06-30

func findLHS(nums []int) int {
	var mapping = make(map[int]int)
	for _, num := range nums {
		mapping[num]++
	}
	ans := 0
	for _, num := range nums {
		if mapping[num+1] > 0 && mapping[num]+mapping[num+1] > ans {
			ans = mapping[num] + mapping[num+1]
		}
	}
	return ans
}
