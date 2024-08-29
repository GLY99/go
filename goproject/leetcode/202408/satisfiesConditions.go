package main

func satisfiesConditions(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i < m-1 && grid[i+1][j] != grid[i][j] {
				return false
			}
			if j < n-1 && grid[i][j+1] == grid[i][j] {
				return false
			}
		}
	}
	return true
}
