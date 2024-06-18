package main

func oddCells(m int, n int, indices [][]int) int {
	rows := make([]int, m)
	cols := make([]int, n)
	for _, p := range indices {
		rows[p[0]]++
		cols[p[1]]++
	}
	ans := 0
	for _, row := range rows {
		for _, col := range cols {
			ans += (row + col) % 2
		}
	}
	return ans
}
