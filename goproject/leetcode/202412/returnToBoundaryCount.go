package main

func returnToBoundaryCount(nums []int) int {
	res := 0
	s := 0
	for _, num := range nums {
		s += num
		if s == 0 {
			res++
		}
	}
	return res
}
