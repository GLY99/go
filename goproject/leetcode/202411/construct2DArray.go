package main

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return make([][]int, 0)
	}
	ans := make([][]int, 0)
	idx := 0
	for i := 0; i < m; i++ {
		ans = append(ans, make([]int, n))
		for j := 0; j < n; j++ {
			ans[i][j] = original[idx]
			idx++
		}
	}
	return ans
}
