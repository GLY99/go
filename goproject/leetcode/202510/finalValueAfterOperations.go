package main

// https://leetcode.cn/problems/final-value-of-variable-after-performing-operations/?envType=daily-question&envId=2025-10-20

func finalValueAfterOperations(operations []string) int {
	var res int = 0
	for _, opoperation := range operations {
		if opoperation == "X++" || opoperation == "++X" {
			res++
		} else {
			res--
		}
	}
	return res
}
