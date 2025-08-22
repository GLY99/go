package main

// https://leetcode.cn/problems/find-the-minimum-area-to-cover-all-ones-i/?envType=daily-question&envId=2025-08-22

func minimumArea(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	minX, minY := m, n
	maxX, maxY := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				minX = min(minX, i)
				maxX = max(maxX, i)
				minY = min(minY, j)
				maxY = max(maxY, j)
			}
		}
	}
	return (maxX - minX + 1) * (maxY - minY + 1)
}
