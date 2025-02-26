package main

// https://leetcode.cn/problems/largest-substring-between-two-equal-characters/?envType=problem-list-v2&envId=hash-table

func maxLengthBetweenEqualCharacters(s string) (res int) {
	var mapping map[rune][]int = map[rune][]int{}
	res = -1
	for idx, c := range s {
		mapping[c] = append(mapping[c], idx)
		if len(mapping[c]) > 1 {
			res = max(mapping[c][len(mapping[c])-1]-mapping[c][0]-1, res)
		}
	}
	return
}
