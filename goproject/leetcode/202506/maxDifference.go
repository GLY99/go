package main

// https://leetcode.cn/problems/maximum-difference-between-even-and-odd-frequency-i/?envType=daily-question&envId=2025-06-10

func maxDifference(s string) int {
	mapping := make([]int, 26)
	for i := 0; i < len(s); i++ {
		mapping[s[i]-'a']++
	}
	maxOdd, minEven := 0, len(s)
	for i := 0; i < 26; i++ {
		if mapping[i] == 0 {
			continue
		}
		if mapping[i]%2 == 1 {
			maxOdd = max(maxOdd, mapping[i])
		} else {
			minEven = min(minEven, mapping[i])
		}
	}
	return maxOdd - minEven
}
