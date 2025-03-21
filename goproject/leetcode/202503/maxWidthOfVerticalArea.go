package main

import "sort"

// https://leetcode.cn/problems/widest-vertical-area-between-two-points-containing-no-points/?envType=problem-list-v2&envId=sorting

func maxWidthOfVerticalArea(points [][]int) int {
	if len(points) < 2 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool { return points[i][0] < points[j][0] })
	ans := 0
	for i := 1; i < len(points); i++ {
		ans = max(ans, points[i][0]-points[i-1][0])
	}
	return ans
}
