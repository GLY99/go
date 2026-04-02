package main

import "math"

// https://leetcode.cn/problems/max-pair-sum-in-an-array/?envType=problem-list-v2&envId=array
func maxSum(nums []int) int {
	ans := -1
	// 一个map,记录位数中存在0-9的数对应的最大数
	maxNum := [10]int{}
	for i := 0; i < 10; i++ {
		maxNum[i] = math.MinInt
	}
	for _, num := range nums {
		// 当前数，位数的最大值
		maxD := 0
		t := num
		for t > 0 {
			maxD = max(maxD, t%10)
			t /= 10
		}

		// 更新最大值
		ans = max(ans, num+maxNum[maxD])
		// 记录当前位数最大值对应的数
		maxNum[maxD] = max(maxNum[maxD], num)
	}
	return ans
}
