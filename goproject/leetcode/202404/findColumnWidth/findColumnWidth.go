package main

func findColumnWidth(grid [][]int) []int {
	res := make([]int, len(grid[0]))
	m := len(grid)
	for i := 0; i < m; i++ {
		n := len(grid[i])
		for j := 0; j < n; j++ {
			x, length := grid[i][j], 0
			if x <= 0 {
				length = 1
			}
			for x != 0 {
				x = x / 10
				length++
			}
			res[j] = max(res[j], length)
		}
	}
	return res
}

func main() {
	findColumnWidth([][]int{{-12}})
}
