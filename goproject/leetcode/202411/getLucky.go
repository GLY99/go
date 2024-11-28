package main

import (
	"strconv"
)

func getLucky(s string, k int) int {
	strNum := ""
	for i := 0; i < len(s); i++ {
		strNum += strconv.Itoa(int(s[i] - 'a' + 1))
	}
	for i := 0; i < k; i++ {
		sum := 0
		for i := 0; i < len(strNum); i++ {
			sum += int(strNum[i] - '0')
		}
		strNum = strconv.Itoa(sum)
		if len(strNum) < 2 {
			break
		}
	}
	ans, _ := strconv.Atoi(strNum)
	return ans
}
