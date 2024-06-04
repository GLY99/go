package main

func projectionArea(grid [][]int) int {
	var xyArea, yzArea, zxArea int
	for i, row := range grid {
		rowMax, colMax := 0, 0
		for j, v := range row {
			if v > 0 {
				xyArea++
			}
			rowMax = max(rowMax, v)
			colMax = max(colMax, grid[j][i])
		}
		yzArea += colMax
		zxArea += rowMax
	}
	return xyArea + yzArea + zxArea
}
