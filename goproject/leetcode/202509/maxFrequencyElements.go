package main

// https://leetcode.cn/problems/count-elements-with-maximum-frequency/?envType=daily-question&envId=2025-09-22

func maxFrequencyElements(nums []int) int {
	maxFreq := 0
	mapping := make(map[int]int)
	for _, num := range nums {
		mapping[num]++
		maxFreq = max(mapping[num], maxFreq)
	}
	count := 0
	for _, v := range mapping {
		if v == maxFreq {
			count++
		}
	}
	return count * maxFreq
}
