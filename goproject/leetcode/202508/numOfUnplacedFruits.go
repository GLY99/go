package main

// https://leetcode.cn/problems/fruits-into-baskets-ii/?envType=daily-question&envId=2025-08-05

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	ans := 0
	for i := 0; i < len(fruits); i++ {
		tmp := 1
		for j := 0; j < len(baskets); j++ {
			if fruits[i] <= baskets[j] {
				tmp = 0
				baskets[j] = 0
				break
			}
		}
		ans += tmp
	}
	return ans
}
