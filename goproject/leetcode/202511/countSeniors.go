package main

// https://leetcode.cn/problems/number-of-senior-citizens/?envType=problem-list-v2&envId=array

func countSeniors(details []string) int {
	counter := 0
	for _, detail := range details {
		ageStr := detail[11:13]
		if ageStr[0] == '6' && ageStr[1] >= '1' || ageStr[0] > '6' {
			counter++
		}
	}
	return counter
}
