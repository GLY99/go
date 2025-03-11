package main

// https://leetcode.cn/problems/check-if-all-characters-have-equal-number-of-occurrences/?envType=problem-list-v2&envId=hash-table

func areOccurrencesEqual(s string) bool {
	mapping := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mapping[s[i]]++
	}
	preCount := mapping[s[0]]
	for k, v := range mapping {
		if k == s[0] {
			continue
		}
		if v != preCount {
			return false
		} else {
			preCount = v
		}
	}
	return true
}
