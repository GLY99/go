package main

// https://leetcode.cn/problems/sum-of-all-odd-length-subarrays/?envType=problem-list-v2&envId=array

func sumOddLengthSubarrays(arr []int) int {
	ans := 0
	for start := range arr {
		for length := 1; start+length < len(arr); length += 2 {
			for _, v := range arr[start : start+length-1] {
				ans += v
			}
		}
	}
	return ans
}

func sumOddLengthSubarrays1(arr []int) int {
	preSum := make([]int, len(arr)+1)
	for i, v := range arr {
		preSum[i+1] = preSum[i] + v
	}
	sum := 0
	for start := range arr {
		for l := 1; start+l <= len(arr); l += 2 {
			end := start + l - 1
			sum += preSum[end+1] - preSum[start]
		}
	}
	return sum
}
