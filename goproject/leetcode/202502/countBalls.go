package main

// https://leetcode.cn/problems/maximum-number-of-balls-in-a-box/description/?envType=daily-question&envId=2025-02-13
func countBalls(lowLimit int, highLimit int) int {
	count := make(map[int]int)
	ans := 0
	for i := lowLimit; i <= highLimit; i++ {
		sum := 0
		for x := i; x > 0; x /= 10 {
			sum += x % 10
		}
		count[sum]++
		ans = max(count[sum], ans)
	}
	return ans
}
