package main

import "strconv"

func calPoints(operations []string) int {
	ans := 0
	points := make([]int, 0)
	for _, op := range operations {
		length := len(points)
		switch op {
		case "+":
			ans += points[length-1] + points[length-2]
			points = append(points, points[length-1]+points[length-2])
		case "D":
			ans += points[length-1] * 2
			points = append(points, points[length-1]*2)
		case "C":
			ans -= points[length-1]
			points = points[:length-1]
		default:
			n, _ := strconv.Atoi(op)
			ans += n
			points = append(points, n)
		}
	}
	return ans
}
