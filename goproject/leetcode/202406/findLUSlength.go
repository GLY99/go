package main

func findLUSlength(a string, b string) int {
	if a != b {
		return max(len(a), len(b))
	}
	return -1
}
