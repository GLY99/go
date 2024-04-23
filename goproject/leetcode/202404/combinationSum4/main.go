package main

import "fmt"

func combinationSum4(nums []int, target int) int {
	// dp[x] 和为x的组合个数
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func main() {
	nums := []int{1, 2, 3}
	target := 7
	fmt.Println(combinationSum4(nums, target))
}
