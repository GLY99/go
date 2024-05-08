package main

import (
	"sort"
)

func findLHS1(nums []int) int {
	sort.Ints(nums)
	start := 0
	ans := 0
	for idx, num := range nums {
		for num-nums[start] > 1 {
			start++
		}
		if num-nums[start] == 1 && idx-start+1 > ans {
			ans = idx - start + 1
		}
	}
	return ans
}

func findLHS2(nums []int) int {
	var myMap map[int]int
	myMap = map[int]int{}
	for _, num := range nums {
		myMap[num]++
	}
	ans := 0
	for _, num := range nums {
		if myMap[num+1] > 0 && myMap[num]+myMap[num+1] > ans {
			ans = myMap[num] + myMap[num+1]
		}
	}
	return ans
}

func main() {
	findLHS1([]int{1, 2})
	findLHS2([]int{1, 2})
}
