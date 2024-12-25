package main

func areSimilar(mat [][]int, k int) bool {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] != mat[i][(j+k)%len(mat[i])] {
				return false
			}
		}
	}
	return true
}
