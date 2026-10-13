package main

// https://leetcode.cn/problems/check-whether-two-strings-are-almost-equivalent/?envType=problem-list-v2&envId=hash-table

func checkAlmostEquivalent(word1 string, word2 string) bool {
	w1Mapping := [26]int{}
	w2Mapping := [26]int{}
	n1 := len(word1)
	n2 := len(word2)
	maxLength := max(n1, n2)
	for i := 0; i < maxLength; i++ {
		if i < n1 {
			w1Mapping[word1[i]-'a']++
		}
		if i < n2 {
			w2Mapping[word2[i]-'a']++
		}
	}
	for i := 0; i < 26; i++ {
		if w1Mapping[i]-w2Mapping[i] > 3 || w1Mapping[i]-w2Mapping[i] < -3 {
			return false
		}
	}
	return true
}
