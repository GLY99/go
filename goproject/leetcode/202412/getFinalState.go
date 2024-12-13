package main

//https://leetcode.cn/problems/final-array-state-after-k-multiplication-operations-i/?envType=daily-question&envId=2024-12-13

func getFinalState(nums []int, k int, multiplier int) []int {
	for i := 0; i < k; i++ {
		l := 0
		for j, _ := range nums {
			if nums[j] < nums[l] {
				l = j
			}
		}
		nums[l] *= multiplier
	}
	return nums
}
