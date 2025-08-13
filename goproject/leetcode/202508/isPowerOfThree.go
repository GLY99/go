package main

// https://leetcode.cn/problems/power-of-three/?envType=daily-question&envId=2025-08-13

func isPowerOfThree(n int) bool {
	for n > 0 && n%3 == 0 {
		n /= 3
	}
	return n == 1
}
