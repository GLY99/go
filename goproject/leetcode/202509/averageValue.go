package main

// https://leetcode.cn/problems/average-value-of-even-numbers-that-are-divisible-by-three/?envType=problem-list-v2&envId=array

func averageValue(nums []int) int {
	sum := 0
	count := 0
	for _, num := range nums {
		if num%3 == 0 && num%2 == 0 {
			sum += num
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return sum / count
}
