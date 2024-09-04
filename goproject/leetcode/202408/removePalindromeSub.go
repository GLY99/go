package main

func removePalindromeSub(s string) int {
	var start int = 0
	var end int = len(s) - 1
	for start < end {
		if s[start] != s[end] {
			return 2
		}
		start++
		end--
	}
	return 1
}
