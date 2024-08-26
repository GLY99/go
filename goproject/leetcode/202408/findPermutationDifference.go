package main

func findPermutationDifference(s string, t string) int {
	var sMapping map[byte]int
	var tMapping map[byte]int
	sMapping = make(map[byte]int, 0)
	tMapping = make(map[byte]int, 0)
	ans := 0
	for i := 0; i < len(s); i++ {
		sMapping[s[i]] = i
		tMapping[t[i]] = i
		idx1, ok1 := sMapping[s[i]]
		idx2, ok2 := tMapping[s[i]]
		if ok1 && ok2 {
			ans += abs(idx1, idx2)
		}
		idx1, ok1 = sMapping[t[i]]
		idx2, ok2 = tMapping[t[i]]
		if ok1 && ok2 {
			ans += abs(idx1, idx2)
		}
	}
	return ans
}
