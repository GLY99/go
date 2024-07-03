package main

import "sort"

func arrayRankTransform(arr []int) []int {
	tmpArr := append([]int{}, arr...)
	sort.Ints(tmpArr)
	mapping := make(map[int]int)
	for idx, _ := range tmpArr {
		if _, ok := mapping[tmpArr[idx]]; !ok {
			mapping[tmpArr[idx]] = len(mapping) + 1
		}
	}
	ans := []int{}
	for idx, _ := range arr {
		ans = append(ans, mapping[arr[idx]])
	}
	return ans
}
