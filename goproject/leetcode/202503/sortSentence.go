package main

import (
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	length := len(strings.Split(s, " "))
	words := make([]string, length)
	for _, word := range strings.Split(s, " ") {
		idx, _ := strconv.Atoi(string(word[len(word)-1]))
		words[idx] = word[:len(word)-1]
	}
	ans := ""
	for idx, word := range words {
		if idx == 0 {
			ans += word
		} else {
			ans += " " + word
		}
	}
	return ans
}
