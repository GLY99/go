package main

import "sort"

func deleteGreatestValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		sort.Ints(grid[i])
	}
	res := 0
	for i := 0; i < n; i++ {
		tmp := 0
		for j := 0; j < m; j++ {
			tmp = max(tmp, grid[j][i])
		}
		res += tmp
	}
	return res
}
