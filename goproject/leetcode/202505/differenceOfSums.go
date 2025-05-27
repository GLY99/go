package main

// https://leetcode.cn/problems/divisible-and-non-divisible-sums-difference/?envType=daily-question&envId=2025-05-27

func differenceOfSums(n int, m int) int {
	sum1, sum2 := 0, 0
	for i := 1; i <= n; i++ {
		if i%m == 0 {
			sum2 += i
		} else {
			sum1 += i
		}
	}
	return sum1 - sum2
}
