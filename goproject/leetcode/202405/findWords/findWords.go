package main

import "unicode"

// https://leetcode.cn/problems/keyboard-row/

func findWords(words []string) []string {
	const rowIdx = "12210111011122000010020202"
	ans := []string{}
lable1:
	for _, word := range words {
		idx := rowIdx[unicode.ToLower(rune(word[0]))-'a']
		for _, c := range word[1:] {
			if rowIdx[unicode.ToLower(rune(c))-'a'] != idx {
				continue lable1
			}
		}
		ans = append(ans, word)
	}
	return ans
}

func main() {
	findWords([]string{"abc"})
}
