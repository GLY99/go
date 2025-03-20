package main

import "sort"

// https://leetcode.cn/problems/mean-of-array-after-removing-some-elements/?envType=problem-list-v2&envId=sorting

func trimMean(arr []int) float64 {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	length := len(arr)
	sum := 0
	for _, num := range arr[length/20 : 19*length/20] {
		sum += num
	}
	return float64(sum) / float64(float64(length)*float64(0.9))
}
