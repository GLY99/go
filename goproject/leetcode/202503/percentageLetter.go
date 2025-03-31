package main

// https://leetcode.cn/problems/percentage-of-letter-in-string/?envType=daily-question&envId=2025-03-31
func percentageLetter(s string, letter byte) int {
	mapping := make(map[byte]int)
	for _, c := range s {
		mapping[byte(c)]++
	}
	return 100 * mapping[letter] / len(s)
}
