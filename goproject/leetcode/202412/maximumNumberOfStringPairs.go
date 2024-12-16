package main

func isReserve(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[len(s2)-1-i] {
			return false
		}
	}
	return true
}

func maximumNumberOfStringPairs(words []string) int {
	ans := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isReserve(words[i], words[j]) {
				ans++
			}
		}
	}
	return ans
}
