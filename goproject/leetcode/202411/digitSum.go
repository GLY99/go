package main

import "strconv"

func digitSum(s string, k int) string {
	for len(s) > k {
		newStr := ""
		for i := 0; i < len(s); i += k {
			sum := 0
			for j := i; j < min(i+k, len(s)); j++ {
				num, _ := strconv.Atoi(string(s[j]))
				sum += num
			}
			newStr += strconv.Itoa(sum)
		}
		s = newStr
	}
	return s
}
