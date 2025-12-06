package main

// https://leetcode.cn/problems/sum-of-squares-of-special-elements/?envType=problem-list-v2&envId=array

func sumOfSquares(nums []int) int {
	sum := 0
	n := len(nums)
	for idx, num := range nums {
		if n%(idx+1) == 0 {
			sum += num * num
		}
	}
	return sum
}
