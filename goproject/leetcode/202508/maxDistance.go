package main

// https://leetcode.cn/problems/two-furthest-houses-with-different-colors/?envType=problem-list-v2&envId=array

func maxDistance(colors []int) int {
	n := len(colors)
	if colors[0] != colors[n-1] {
		return n - 1
	}
	i, j := 1, n-2
	for ; i < n; i++ {
		if colors[i] != colors[0] {
			break
		}
	}
	for ; j > 0; j-- {
		if colors[j] != colors[n-1] {
			break
		}
	}
	return max(j, n-i-1)
}
