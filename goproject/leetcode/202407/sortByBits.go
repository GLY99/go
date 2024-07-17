package main

import "sort"

func onesCount(num int) int {
	count := 0
	for num > 0 {
		count += num % 2
		num = num / 2
	}
	return count
}

func sortByBits(arr []int) []int {
	sort.Slice(
		arr, func(i, j int) bool {
			c1 := onesCount(arr[i])
			c2 := onesCount(arr[j])
			if c1 == c2 {
				return arr[i] < arr[j]
			} else {
				return c1 < c2
			}
		},
	)
	return arr
}
