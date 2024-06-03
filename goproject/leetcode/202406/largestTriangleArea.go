package main

import "math"

func triangleArea(point1 []int, point2 []int, point3 []int) (ans float64) {
	x1, y1 := point1[0], point1[1]
	x2, y2 := point2[0], point2[1]
	x3, y3 := point3[0], point3[1]
	ans = math.Abs(float64(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2))) * 0.5
	return
}

func largestTriangleArea(points [][]int) float64 {
	var res float64 = 0.0
	for i, point1 := range points {
		for j, point2 := range points[:i] {
			for _, point3 := range points[:j] {
				res = math.Max(res, triangleArea(point1, point2, point3))
			}
		}
	}
	return res
}
