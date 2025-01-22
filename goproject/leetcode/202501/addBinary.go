package main

import "strconv"

func addBinary(a string, b string) string {
	if a[0] == '0' {
		return b
	} else if b[0] == '0' {
		return a
	}
	aLen := len(a)
	bLen := len(b)
	maxLen := max(aLen, bLen)
	carry := 0
	ans := ""
	for i := 0; i < maxLen; i++ {
		if i < aLen {
			carry += int(a[aLen-i-1] - '0')
		}
		if i < bLen {
			carry += int(b[bLen-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry = carry / 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}
