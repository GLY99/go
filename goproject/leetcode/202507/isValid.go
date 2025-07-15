package main

import "unicode"

// https://leetcode.cn/problems/valid-word/?envType=daily-question&envId=2025-07-15

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	hasVowel := false
	hasConsonant := false
	for _, c := range word {
		if unicode.IsLetter(c) {
			ch := unicode.ToLower(c)
			if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
				hasVowel = true
			} else {
				hasConsonant = true
			}
		} else if !unicode.IsDigit(c) {
			return false
		}
	}
	return hasVowel && hasConsonant
}
