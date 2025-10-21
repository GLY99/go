package main

// https://leetcode.cn/problems/maximum-count-of-positive-integer-and-negative-integer/?envType=problem-list-v2&envId=array

func maximumCount(nums []int) int {
	counter1, counter2 := 0, 0
	for _, num := range nums {
		if num < 0 {
			counter1++
		} else if num > 0 {
			counter2++
		}
	}
	return max(counter1, counter2)
}
