package main

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			sum, num := 0, 0
			for a := max(i-1, 0); a < min(i+2, m); a++ {
				for b := max(j-1, 0); b < min(j+2, n); b++ {
					sum += img[a][b]
					num++
				}
			}
			ans[i][j] = sum / num
		}
	}
	return ans
}
