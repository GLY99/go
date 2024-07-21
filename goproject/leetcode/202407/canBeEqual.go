package main

import "sort"

func canBeEqual(target []int, arr []int) bool {
	if len(arr) != len(target) {
		return false
	}
	sort.Ints(arr)
	sort.Ints(target)
	for idx := range target {
		if arr[idx] != target[idx] {
			return false
		}
	}
	return true
}
