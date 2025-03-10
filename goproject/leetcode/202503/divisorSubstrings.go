package main

import "strconv"

// https://leetcode.cn/problems/find-the-k-beauty-of-a-number/?envType=daily-question&envId=2025-03-10

func divisorSubstrings(num int, k int) int {
	ans := 0
	numStr := strconv.Itoa(num)
	for i := 0; i < len(numStr)-k+1; i++ {
		subNum, _ := strconv.Atoi(numStr[i : i+k])
		if subNum != 0 && num%subNum == 0 {
			ans++
		}
	}
	return ans
}
