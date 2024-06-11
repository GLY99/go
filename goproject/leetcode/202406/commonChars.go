package main

import (
	"math"
)

func commonChars(words []string) []string {
	minFreq := new([26]int)
	for i, _ := range minFreq {
		minFreq[i] = math.MaxInt64
	}
	for _, word := range words {
		freq := new([26]int)
		for _, c := range word {
			freq[c-'a']++
		}
		for i, f := range freq {
			minFreq[i] = min(f, minFreq[i])
		}
	}
	ans := make([]string, 0)
	for i, f := range minFreq {
		for ; f > 0; f-- {
			ans = append(ans, string(i+'a'))
		}
	}
	return ans
}
