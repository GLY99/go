package main

// https://leetcode.cn/problems/lucky-numbers-in-a-matrix/?envType=problem-list-v2&envId=array

func luckyNumbers(matrix [][]int) []int {
	minRow := make([]int, len(matrix))
	maxCol := make([]int, len(matrix[0]))
	for i, row := range matrix {
		minRow[i] = row[0]
		for j, x := range row {
			minRow[i] = min(minRow[i], x)
			maxCol[j] = max(maxCol[j], x)
		}
	}
	ans := make([]int, 0)
	for i, row := range matrix {
		for j, x := range row {
			if x == minRow[i] && x == maxCol[j] {
				ans = append(ans, x)
			}
		}
	}
	return ans
}
