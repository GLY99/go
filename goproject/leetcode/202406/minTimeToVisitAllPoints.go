package main

import "math"

func minTimeToVisitAllPoints(points [][]int) int {
	ans := 0
	length := len(points)
	for idx, point := range points {
		if idx == length-1 {
			break
		}
		x0, y0 := point[0], point[1]
		x1, y1 := points[idx+1][0], points[idx+1][1]
		ans += int(math.Max(math.Abs(float64(x0-x1)), math.Abs(float64(y0-y1))))
	}
	return ans
}
