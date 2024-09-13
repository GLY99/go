package main

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	var freq map[string]int = make(map[string]int)
	insert := func(s string) {
		sSlice := strings.Split(s, " ")
		for _, word := range sSlice {
			freq[word]++
		}
	}
	insert(s1)
	insert(s2)
	var ans []string
	for word, count := range freq {
		if count == 1 {
			ans = append(ans, word)
		}
	}
	return ans
}
