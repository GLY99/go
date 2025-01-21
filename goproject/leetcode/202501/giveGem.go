package main

import "sort"

func giveGem(gem []int, operations [][]int) int {
	for _, operation := range operations {
		gem[operation[1]] += gem[operation[0]] / 2
		gem[operation[0]] -= gem[operation[0]] / 2
	}
	sort.Slice(gem, func(i, j int) bool { return gem[i] < gem[j] })
	return gem[0] - gem[len(gem)-1]
}
