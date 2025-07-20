package main

// https://leetcode.cn/problems/sign-of-the-product-of-an-array/?envType=problem-list-v2&envId=array

func arraySign(nums []int) int {
	count := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		} else if num < 0 {
			count++
		}
	}
	if count%2 == 0 {
		return 1
	} else {
		return -1
	}
}
