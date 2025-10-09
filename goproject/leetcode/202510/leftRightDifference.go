package main

// https://leetcode.cn/problems/left-and-right-sum-differences/?envType=problem-list-v2&envId=array

func leftRightDifference(nums []int) []int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	leftSum := 0
	rightSum := 0
	ans := make([]int, len(nums))
	for idx, num := range nums {
		rightSum = sum - num - leftSum
		ans[idx] = abs(leftSum - rightSum)
		leftSum += num
	}
	return ans
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
