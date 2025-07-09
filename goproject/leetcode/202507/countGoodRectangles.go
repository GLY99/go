package main

// https://leetcode.cn/problems/number-of-rectangles-that-can-form-the-largest-square/?envType=problem-list-v2&envId=array

func countGoodRectangles(rectangles [][]int) int {
	mapping := make(map[int]int)
	maxLength := 0
	for _, v := range rectangles {
		minVal := min(v[0], v[1])
		mapping[minVal]++
		maxLength = max(minVal, maxLength)
	}
	return mapping[maxLength]
}
