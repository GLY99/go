package main

func surfaceArea(grid [][]int) int {
	if grid == nil || len(grid) < 1 {
		return 0
	}
	count := 0
	covers := 0
	for i, row := range grid {
		for j, v := range row {
			count += v
			// 统计当前位置由于堆叠导致覆盖的面
			if v > 0 {
				covers += v - 1
			}
			// 统计当前位置上一行同一列覆盖了几个面
			if i > 0 {
				covers += min(v, grid[i-1][j])
			}
			// 统计当前位置上一列同一行覆盖了几个面
			if j > 0 {
				covers += min(v, grid[i][j-1])
			}
		}
	}
	return count*6 - covers*2
}
