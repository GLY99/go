package main

func repeatedNTimes(nums []int) int {
	mapping := make(map[int]int)
	length := len(nums)
	for _, num := range nums {
		mapping[num]++
	}
	for k, v := range mapping {
		if v == length/2 {
			return k
		}
	}
	return -1
}
