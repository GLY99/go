package main

func countNegatives(grid [][]int) int {
	rowLength := len(grid)
	colLength := len(grid[0])
	count := 0
	for i := 0; i < rowLength; i++ {
		left := 0
		right := colLength - 1
		for left <= right {
			mid := left + (right-left)/2
			if grid[i][mid] >= 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		count += colLength - left
	}
	return count
}
