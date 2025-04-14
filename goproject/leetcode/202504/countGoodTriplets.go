package main

// https://leetcode.cn/problems/count-good-triplets/?envType=daily-question&envId=2025-04-14

func countGoodTriplets(arr []int, a int, b int, c int) int {
	length := len(arr)
	count := 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					count++
				}
			}
		}
	}
	return count
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
