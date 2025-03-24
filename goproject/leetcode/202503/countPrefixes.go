package main

// https://leetcode.cn/problems/count-prefixes-of-a-given-string/?envType=daily-question&envId=2025-03-24

func countPrefixes(words []string, s string) int {
	ans := 0
	for _, word := range words {
		if len(word) > len(s) {
			continue
		}
		flag := true
		for i := 0; i < len(word); i++ {
			if word[i] != s[i] {
				flag = false
				break
			}
		}
		if flag {
			ans++
		}
	}
	return ans
}
