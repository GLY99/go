package main

// https://leetcode.cn/problems/count-common-words-with-one-occurrence/?envType=problem-list-v2&envId=array

func countWords(words1 []string, words2 []string) int {
	words1Map := make(map[string]int)
	words2Map := make(map[string]int)
	for _, word := range words1 {
		words1Map[word]++
	}
	for _, word := range words2 {
		words2Map[word]++
	}
	ans := 0
	for word, count := range words1Map {
		if count == 1 && words2Map[word] == 1 {
			ans++
		}
	}
	return ans
}
