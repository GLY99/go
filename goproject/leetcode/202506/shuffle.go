package main

// https://leetcode.cn/problems/shuffle-the-array/?envType=problem-list-v2&envId=array

func shuffle(nums []int, n int) []int {
	ans := make([]int, 2*n)
	idx := 0
	for i := 0; i < 2*n; i += 2 {
		ans[i] = nums[idx]
		ans[i+1] = nums[n+idx]
		idx++
	}
	return ans
}
