package main

// https://leetcode.cn/problems/find-words-containing-character/?envType=daily-question&envId=2025-05-24

func findWordsContaining(words []string, x byte) []int {
	ans := make([]int, 0)
	for idx, word := range words {
		for _, char := range word {
			if char == rune(x) {
				ans = append(ans, idx)
				break
			}
		}
	}
	return ans
}
