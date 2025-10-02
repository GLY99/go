package main

// https://leetcode.cn/problems/water-bottles/?envType=daily-question&envId=2025-10-01

func numWaterBottles(numBottles int, numExchange int) int {
	ans := numBottles
	bottles := numBottles
	for bottles >= numExchange {
		newBottles := bottles / numExchange
		ans += newBottles
		bottles -= newBottles * numExchange
		bottles += newBottles
	}
	return ans
}
