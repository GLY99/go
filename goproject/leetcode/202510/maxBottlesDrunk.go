package main

// https://leetcode.cn/problems/water-bottles-ii/?envType=daily-question&envId=2025-10-01

func maxBottlesDrunk(numBottles int, numExchange int) int {
	ans := numBottles
	// 只要i >= numExchange就能换一瓶； 换完后i减少了numExchange - 1个
	for i := numBottles; i >= numExchange; numExchange++ {
		ans++
		i -= numExchange - 1
	}
	return ans
}
