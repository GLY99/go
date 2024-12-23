package main

func circularGameLosers(n int, k int) []int {
	visit := make(map[int]bool, n)
	j := 0
	for i := k; !visit[j]; i += k {
		visit[j] = true
		j = (j + i) % n
	}
	ans := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if !visit[i] {
			ans = append(ans, i+1)
		}
	}
	return ans
}
