package main

// https://leetcode.cn/problems/redistribute-characters-to-make-all-strings-equal/?envType=problem-list-v2&envId=hash-table

func makeEqual(words []string) bool {
	mapping := make([]int, 26)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			mapping[word[i]-'a']++
		}
	}
	for i := 0; i < len(mapping); i++ {
		if mapping[i]%len(words) != 0 {
			return false
		}
	}
	return true
}
