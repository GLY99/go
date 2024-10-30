package main

func getSmallestString(s string) string {
	r := []rune(s)
	for i := 0; i < len(r)-1; i++ {
		if r[i] > r[i+1] && r[i]%2 == r[i+1]%2 {
			r[i], r[i+1] = r[i+1], r[i]
			break
		}
	}
	return string(r)
}
