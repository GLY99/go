package main

import "strings"

// https://leetcode.cn/problems/string-matching-in-an-array/?envType=problem-list-v2&envId=array

func stringMatching(words []string) []string {
	ans := []string{}
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i != j && strings.Contains(words[j], words[i]) {
				ans = append(ans, words[j])
				break
			}
		}
	}
	return ans
}
