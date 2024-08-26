package main

func maximumLengthSubstring(s string) int {
	length := len(s)
	l := 0
	r := 0
	mapping := map[byte]int{}
	ans := 0
	for r < length {
		for mapping[s[r]] >= 2 {
			mapping[s[l]]--
			l++
		}
		mapping[s[r]]++
		r++
		ans = max(ans, r-l)
	}
	return ans
}
