package main

// https://leetcode.cn/problems/rings-and-rods/?envType=problem-list-v2&envId=hash-table

func countPoints(rings string) int {
	mapping := [10][3]int{}
	n := len(rings)
	cMapping := map[byte]int{'R': 0, 'G': 1, 'B': 2}
	set := make(map[int]struct{})
	for i := 0; i < n-1; i += 2 {
		mapping[rings[i+1]-'0'][cMapping[rings[i]]] += 1
		if mapping[rings[i+1]-'0'][0] > 0 && mapping[rings[i+1]-'0'][1] > 0 && mapping[rings[i+1]-'0'][2] > 0 {
			set[int(rings[i+1]-'0')] = struct{}{}
		}
	}
	return len(set)
}
