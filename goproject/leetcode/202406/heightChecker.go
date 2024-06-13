package main

import "sort"

func heightChecker(heights []int) int {
	sroted := append([]int{}, heights...)
	sort.Ints(sroted)
	res := 0
	for i, v := range heights {
		if v != sroted[i] {
			res++
		}
	}
	return res
}
