package main

import "sort"

// https://leetcode.cn/problems/sort-matrix-by-diagonals/?envType=daily-question&envId=2025-08-28

func sortMatrix(grid [][]int) [][]int {
	length := len(grid)
	// 左下三角形
	// 每次列从0开始增加
	// 00 11 22 行 = 列 + 0
	// 10 21 行 = 列 + 1
	// 20 行 = 列 + 2
	for i := 0; i < length; i++ {
		tmp := []int{}
		for j := 0; i+j < length; j++ {
			tmp = append(tmp, grid[i+j][j])
		}
		sort.Slice(tmp, func(i, j int) bool { return tmp[i] > tmp[j] })
		for j := 0; i+j < length; j++ {
			grid[i+j][j] = tmp[j]
		}
	}
	// 右上三角形，不含对角线
	// 行每次都是从0开始
	// 01 12 列 = 行 + 1
	// 02 列 = 行 + 2
	for j := 1; j < length; j++ {
		tmp := []int{}
		for i := 0; i+j < length; i++ {
			tmp = append(tmp, grid[i][i+j])
		}
		sort.Slice(tmp, func(i, j int) bool { return tmp[i] < tmp[j] })
		for i := 0; i+j < length; i++ {
			grid[i][i+j] = tmp[i]
		}
	}
	return grid
}
