package main

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	bestMin := math.MaxInt32
	length := len(arr)
	ans := make([][]int, 0)
	for i := 0; i < length-1; i++ {
		delta := arr[i+1] - arr[i]
		if delta < bestMin {
			bestMin = delta
			ans = [][]int{{arr[i], arr[i+1]}}
		} else if delta == bestMin {
			ans = append(ans, []int{arr[i], arr[i+1]})
		}
	}
	return ans
}
