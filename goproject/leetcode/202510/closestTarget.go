package main

// https://leetcode.cn/problems/shortest-distance-to-target-string-in-a-circular-array/?envType=problem-list-v2&envId=array

func closestTarget(words []string, target string, startIndex int) int {
	l, r := startIndex, startIndex
	n := len(words)
	for i := 0; i < n; i++ {
		if words[l] == target || words[r] == target {
			return i
		}
		l = (l - 1 + n) % n
		r = (r + 1) % n
	}
	return -1
}
