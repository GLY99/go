package main

import "sort"

// https://leetcode.cn/problems/apple-redistribution-into-boxes/?envType=problem-list-v2&envId=sorting

func minimumBoxes(apple []int, capacity []int) int {
	sum := 0
	for i := 0; i < len(apple); i++ {
		sum += apple[i]
	}
	sort.Slice(capacity, func(i, j int) bool { return capacity[i] > capacity[j] })
	count := 0
	for _, c := range capacity {
		sum -= c
		count++
		if sum <= 0 {
			break
		}
	}
	return count
}
