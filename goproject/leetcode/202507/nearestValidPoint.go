package main

// https://leetcode.cn/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/?envType=problem-list-v2&envId=array

func nearestValidPoint(x int, y int, points [][]int) int {
	ans := -1
	minLength := -1
	for idx, point := range points {
		if point[0] != x && point[1] != y {
			continue
		}
		l := abs(point[0]-x) + abs(point[1]-y)
		if minLength == -1 || l < minLength {
			minLength = l
			ans = idx
		}
	}
	return ans
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
