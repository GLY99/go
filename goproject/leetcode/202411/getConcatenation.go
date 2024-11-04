package main

func getConcatenation(nums []int) []int {
	ans := make([]int, 0)
	ans = append(ans, nums...)
	ans = append(ans, nums...)
	return ans
}
