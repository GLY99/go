package main

func zigzagTraversal(grid [][]int) []int {
	ans := make([]int, 0)
	ok := true
	for i := 0; i < len(grid); i++ {
		if i%2 == 0 {
			for j := 0; j < len(grid[i]); j++ {
				if ok {
					ans = append(ans, grid[i][j])
				}
				ok = !ok
			}
		} else {
			for j := len(grid[i]) - 1; j >= 0; j-- {
				if ok {
					ans = append(ans, grid[i][j])
				}
				ok = !ok
			}
		}
	}
	return ans
}
