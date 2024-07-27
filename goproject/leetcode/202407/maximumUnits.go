package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool { return boxTypes[i][1] > boxTypes[j][1] })
	ans := 0
	for _, p := range boxTypes {
		if p[0] >= truckSize {
			ans += truckSize * p[1]
			break
		}
		ans += p[0] * p[1]
		truckSize -= p[0]
	}
	return ans
}
