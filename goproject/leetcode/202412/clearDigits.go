package main

import "unicode"

// https://leetcode.cn/problems/clear-digits/?envType=problem-list-v2&envId=simulation&status=TO_DO&difficulty=EASY

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func clearDigits1(s string) string {
	newStr := ""
	digitCount := 0
	for i := len(s) - 1; i >= 0; i-- {
		b := s[i]
		if b >= 'a' && b <= 'z' && digitCount == 0 {
			newStr += string(b)
		} else if b >= 'a' && b <= 'z' && digitCount != 0 {
			digitCount--
		} else {
			digitCount++
		}
	}
	return reverseString(newStr)
}

func clearDigits(s string) string {
	var stack []byte
	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, byte(c))
		} else {
			if unicode.IsDigit(c) {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, byte(c))
			}
		}
	}
	return string(stack)
}
