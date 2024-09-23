package main

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	var mapping map[int]int = map[int]int{}
	for _, item := range items1 {
		mapping[item[0]] += item[1]
	}
	for _, item := range items2 {
		mapping[item[0]] += item[1]
	}
	var res [][]int
	for v, w := range mapping {
		res = append(res, []int{v, w})
	}
	sort.Slice(res, func(i, j int) bool { return res[i][0] < res[j][0] })
	return res
}
