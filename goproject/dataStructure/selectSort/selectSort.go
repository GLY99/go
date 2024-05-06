package main

import "fmt"

func selectSort(nums []int) {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		minIdx := i
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			nums[i], nums[minIdx] = nums[minIdx], nums[i]
		}
	}
}

func main() {
	nums := []int{2, 3, 1, 8, 9, 4, 5, 4}
	selectSort(nums)
	fmt.Println(nums)
}
