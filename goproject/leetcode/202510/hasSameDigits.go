package main

// https://leetcode.cn/problems/check-if-digits-are-equal-in-string-after-operations-i/?envType=daily-question&envId=2025-10-23

func hasSameDigits(s string) bool {
	n := len(s)
	if n == 2 {
		return s[0] == s[1]
	}
	slice := []byte(s)
	for i := 1; i <= n-2; i++ {
		for j := 0; j <= n-i-1; j++ {
			slice[j] = byte(((int(slice[j]-'0') + int(slice[j+1]-'0')) % 10) + '0')
		}
	}
	return slice[0] == slice[1]
}
