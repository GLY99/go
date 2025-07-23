package main

// https://leetcode.cn/problems/maximum-population-year/?envType=problem-list-v2&envId=array

func maximumPopulation(logs [][]int) int {
	delta := make([]int, 101)
	offset := 1950
	for _, log := range logs {
		delta[log[0]-offset]++
		delta[log[1]-offset]--
	}
	maxPopulation := 0
	res := 0
	currPopulation := 0
	for idx, d := range delta {
		currPopulation += d
		if currPopulation > maxPopulation {
			maxPopulation = currPopulation
			res = idx
		}
	}
	return res + offset
}
