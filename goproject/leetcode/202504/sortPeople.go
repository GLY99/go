package main

import "sort"

// https://leetcode.cn/problems/sort-the-people/description/?envType=problem-list-v2&envId=sorting

func sortPeople(names []string, heights []int) []string {
	length := len(names)
	idxs := make([]int, length)
	for i := 0; i < length; i++ {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool { return heights[idxs[i]] > heights[idxs[j]] })
	newNames := make([]string, length)
	for i := 0; i < length; i++ {
		newNames[i] = names[idxs[i]]
	}
	return newNames
}
