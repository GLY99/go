package main

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	for i := 1; i < len(arr)-1; i++ {
		if arr[i]*2 != arr[i-1]+arr[i+1] {
			return false
		}
	}
	return true
}
