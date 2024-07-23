package main

func longestPalindrome(s string) int {
	mapping := make(map[rune]int)
	for _, c := range s {
		mapping[c]++
	}
	ans := 0
	for _, v := range mapping {
		ans += v / 2 * 2
		if ans%2 == 0 && v%2 == 1 {
			ans += 1
		}
	}
	return ans
}
