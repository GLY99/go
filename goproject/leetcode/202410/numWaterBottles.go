package main

func numWaterBottles(numBottles int, numExchange int) int {
	bottles := numBottles
	ans := numBottles
	for bottles >= numExchange {
		newBottles := bottles / numExchange
		ans += newBottles
		bottles -= newBottles * numExchange
		bottles += newBottles
	}
	return ans
}
