package main

// https://leetcode.cn/problems/maximum-number-of-balloons/submissions/599497140/?envType=problem-list-v2&envId=hash-table

func maxNumberOfBalloons(text string) int {
	mapping := make(map[byte]int)
	minCount := len(text)
	for i := 0; i < len(text); i++ {
		mapping[text[i]]++
	}
	for _, k := range "balon" {
		if _, ok := mapping[byte(k)]; !ok {
			return 0
		}
		if k == 'l' || k == 'o' {
			mapping[byte(k)] /= 2
		}
		if k == 'b' || k == 'a' || k == 'l' || k == 'o' || k == 'n' {
			minCount = min(minCount, mapping[byte(k)])
		}
	}
	return minCount
}
