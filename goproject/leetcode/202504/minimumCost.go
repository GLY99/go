package main

import "sort"

// https://leetcode.cn/problems/minimum-cost-of-buying-candies-with-discount/?envType=problem-list-v2&envId=sorting

func minimumCost(cost []int) int {
	sort.Slice(cost, func(i, j int) bool { return cost[i] > cost[j] })
	costs := 0
	for i := 0; i < len(cost); i += 3 {
		if i+2 >= len(cost) {
			for ; i < len(cost); i++ {
				costs += cost[i]
			}
		} else {
			costs += cost[i] + cost[i+1]
		}
	}
	return costs
}
