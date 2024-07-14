package main

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	i := 0
	j := 0
	for j < len(t) {
		if i == len(s) {
			return true
		}
		if t[j] == s[i] {
			i++
		}
		j++
	}
	return i == len(s)
}
