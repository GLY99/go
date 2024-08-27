package main

func reverseString(s []byte) {
	length := len(s)
	l := 0
	r := length - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
