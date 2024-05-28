package main

import "fmt"

func dominantIndex(nums []int) int {
	res := 0
	firstMaxNum := -1
	secondMaxNum := -1
	for idx, num := range nums {
		if num > firstMaxNum {
			secondMaxNum = firstMaxNum
			firstMaxNum = num
			res = idx
		} else if num > secondMaxNum {
			secondMaxNum = num
		}
	}
	if firstMaxNum >= 2*secondMaxNum {
		return res
	}
	return -1
}

func main() {
	fmt.Println(dominantIndex([]int{1, 2, 3, 4}))
}
