package main

func getMaximumGenerated(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	nums := make([]int, n+1)
	nums[1] = 1
	for i := 2; i <= n; i++ {
		nums[i] = nums[i/2] + i%2*nums[i/2+1]
	}
	ans := 0
	for _, num := range nums {
		ans = max(num, ans)
	}
	return ans
}
