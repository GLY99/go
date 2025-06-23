package main

func restoreString(s string, indices []int) string {
	newS := []byte(s)
	for i := 0; i < len(s); i++ {
		newS[indices[i]] = s[i]
	}
	return string(newS)
}
