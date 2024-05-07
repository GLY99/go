package main

import "fmt"

func insertSort(nums []int) {
	length := len(nums)
	for i := 1; i < length; i++ {
		insertVal := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if insertVal < nums[j] {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		if j+1 != i {
			nums[j+1] = insertVal
		}
	}
}

func main() {
	nums := []int{2, 3, 1, 8, 9, 4, 5, 4}
	insertSort(nums)
	fmt.Println(nums)
}
