package main

import "strings"

// https://leetcode.cn/problems/maximum-number-of-words-you-can-type/?envType=problem-list-v2&envId=hash-table

func canBeTypedWords(text string, brokenLetters string) int {
	brokens := new([26]int)
	for _, c := range brokenLetters {
		(*brokens)[c-'a']++
	}
	textArr := strings.Split(text, " ")
	ans := 0
lab1:
	for _, word := range textArr {
		for _, c := range word {
			if (*brokens)[c-'a'] != 0 {
				continue lab1
			}
		}
		ans++
	}
	return ans
}
