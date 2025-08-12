package main

import "strings"

// https://leetcode.cn/problems/maximum-number-of-words-found-in-sentences/?envType=problem-list-v2&envId=array

func mostWordsFound(sentences []string) int {
	ans := len(strings.Split(sentences[0], " "))
	for i := 1; i < len(sentences); i++ {
		ans = max(ans, len(strings.Split(sentences[i], " ")))
	}
	return ans
}
