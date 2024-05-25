package main

import (
	"fmt"
	"sort"
)

func findErrorNums1(nums []int) []int {
	res := make([]int, 2)
	mapping := make(map[int]int)
	for _, num := range nums {
		mapping[num]++
		if mapping[num] > 1 {
			res[0] = num
		}
	}
	for i := 1; i <= len(nums); i++ {
		_, ok := mapping[i]
		if !ok {
			res[1] = i
		}
	}
	return res
}

func findErrorNums2(nums []int) []int {
	res := make([]int, 2)
	sort.Ints(nums)
	length := len(nums)
	if nums[0] != 1 {
		res[1] = 1
	} else if nums[length-1] != length {
		res[1] = length
	}
	for i := 0; i < length-1; i++ {
		if nums[i+1] == nums[i] {
			res[0] = nums[i]
		}
		if nums[i+1]-nums[i] == 2 {
			res[1] = nums[i] + 1
		}
	}
	return res
}

func main() {
	fmt.Println(findErrorNums1([]int{1, 2, 2, 3}))
	fmt.Println(findErrorNums2([]int{1, 2, 2, 3}))
}
