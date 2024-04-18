package main

import "sort"

func findOriginalArray(changed []int) []int {
	sort.Ints(changed)
	myMap := make(map[int]int, len(changed))
	for _, v := range changed {
		myMap[v]++
	}
	res := []int{}
	for _, v := range changed {
		if myMap[v] == 0 {
			continue
		}
		myMap[v]--
		if myMap[v*2] == 0 {
			return []int{}
		}
		myMap[v*2]--
		res = append(res, v)
	}
	return res
}

func main() {
	findOriginalArray([]int{1, 3, 4, 2, 6, 8})
}
