package main

import "strconv"

func divisorSubstrings(num int, k int) int {
	ans := 0
	numStr := strconv.FormatInt(int64(num), 10)
	for i := 0; i < len(numStr)-k+1; i++ {
		subNum, _ := strconv.Atoi(numStr[i : i+k])
		if subNum != 0 && num%subNum == 0 {
			ans++
		}
	}
	return ans
}
