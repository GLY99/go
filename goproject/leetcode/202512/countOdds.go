package main

// https://leetcode.cn/problems/count-odd-numbers-in-an-interval-range/?envType=daily-question&envId=2025-12-07

func countOdds(low int, high int) int {
	ans := 0
	for i := low; i <= high; i++ {
		if i%2 != 0 {
			ans++
		}
	}
	return ans
}
