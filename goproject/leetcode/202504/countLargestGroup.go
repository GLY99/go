package main

// https://leetcode.cn/problems/count-largest-group/?envType=daily-question&envId=2025-04-23

func countLargestGroup(n int) int {
	mapping := make(map[int]int)
	maxCount := 0
	for i := 1; i <= n; i++ {
		sum := 0
		val := i
		for val > 0 {
			sum += val % 10
			val /= 10
		}
		mapping[sum]++
		maxCount = max(maxCount, mapping[sum])
	}
	count := 0
	for _, v := range mapping {
		if v == maxCount {
			count++
		}
	}
	return count
}
