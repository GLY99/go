package main

func countKeyChanges(s string) int {
	count := 0
	for i := 1; i < len(s); i++ {
		if s[i-1]-'a' != s[i]-'a' && s[i-1]-'A' != s[i]-'A' && s[i-1]-'A' != s[i]-'a' && s[i-1]-'a' != s[i]-'A' {
			count++
		}
	}
	return count
}
