package main

func findMissingAndRepeatedValues(grid [][]int) []int {
	length := len(grid)
	arr := make([]int, length*length)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			arr[grid[i][j]-1]++
		}
	}
	res := make([]int, 2)
	res[0] = -1
	res[1] = -1
	for idx, v := range arr {
		if res[0] != -1 && res[1] != -1 {
			break
		}
		if v == 2 {
			res[0] = idx + 1
		} else if v == 0 {
			res[1] = idx + 1
		}
	}
	return res
}
