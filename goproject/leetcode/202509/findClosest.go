package main

// https://leetcode.cn/problems/find-closest-person/?envType=daily-question&envId=2025-09-04

func findClosest(x int, y int, z int) int {
	if abs(x, z) < abs(y, z) {
		return 1
	} else if abs(x, z) > abs(y, z) {
		return 2
	} else {
		return 0
	}
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
