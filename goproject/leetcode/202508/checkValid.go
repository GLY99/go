package main

// https://leetcode.cn/problems/check-if-every-row-and-column-contains-all-numbers/?envType=problem-list-v2&envId=array

func checkValid(matrix [][]int) bool {
	n := len(matrix)

	for i := 0; i < n; i++ {
		mapping := make(map[int]interface{})
		for j := 0; j < n; j++ {
			if _, ok := mapping[matrix[i][j]]; ok {
				return false
			}
			mapping[matrix[i][j]] = struct{}{}
		}
	}
	for i := 0; i < n; i++ {
		mapping := make(map[int]interface{})
		for j := 0; j < n; j++ {
			if _, ok := mapping[matrix[j][i]]; ok {
				return false
			}
			mapping[matrix[j][i]] = struct{}{}
		}
	}
	return true
}
