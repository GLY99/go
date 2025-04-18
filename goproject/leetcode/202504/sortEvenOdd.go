package main

import "sort"

func sortEvenOdd(nums []int) []int {
	even, odd := make([]int, 0), make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			even = append(even, nums[i])
		} else {
			odd = append(odd, nums[i])
		}
	}
	sort.Slice(even, func(i, j int) bool { return even[i] < even[j] })
	sort.Slice(odd, func(i, j int) bool { return odd[i] > odd[j] })
	for i := 0; i < len(even); i++ {
		nums[2*i] = even[i]
	}
	for i := 0; i < len(odd); i++ {
		nums[2*i+1] = odd[i]
	}
	return nums
}
