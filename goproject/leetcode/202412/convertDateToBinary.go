package main

import "strconv"

func binary(x int) string {
	var s []byte
	for ; x != 0; x >>= 1 {
		s = append(s, '0'+byte(x&1)) // 用'0'+可以去除前导0
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func convertDateToBinary(date string) string {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])
	return binary(year) + "-" + binary(month) + "-" + binary(day)
}
