package main

func sumZero(n int) []int {
	ans := make([]int, 0)
	for i := 1; i <= n/2; i++ {
		ans = append(ans, i)
		ans = append(ans, -i)
	}
	if n%2 != 0 {
		ans = append(ans, 0)
	}
	return ans
}
