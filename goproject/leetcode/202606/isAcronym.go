package main

// https://leetcode.cn/problems/check-if-a-string-is-an-acronym-of-words/?envType=problem-list-v2&envId=array

func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	for idx, word := range words {
		if word[0] != s[idx] {
			return false
		}
	}
	return true
}
