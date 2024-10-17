package main

func countBinarySubstrings(s string) int {
	counters := make([]int, 0)
	idx, length := 0, len(s)
	for idx < length {
		c := s[idx]
		count := 0
		for idx < length && s[idx] == c {
			idx++
			count++
		}
		counters = append(counters, count)
	}
	ans := 0
	for i := 1; i < len(counters); i++ {
		ans += min(counters[i], counters[i-1])
	}
	return ans
}
