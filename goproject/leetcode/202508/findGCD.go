package main

import "math"

// https://leetcode.cn/problems/find-greatest-common-divisor-of-array/?envType=problem-list-v2&envId=array

func findGCD(nums []int) int {
	maxNum := 0
	minNum := math.MaxInt
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
		if num < minNum {
			minNum = num
		}
	}
	return gcd(minNum, maxNum)
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		tmp := a % b
		a = b
		b = tmp
	}
	return a
}
