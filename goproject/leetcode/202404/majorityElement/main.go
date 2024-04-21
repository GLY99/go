package main

import "fmt"

func majorityElement(nums []int) int {
	num := 0
	count := 0
	for _, v := range nums {
		if count == 0 {
			num = v
		}
		if v == num {
			count++
		} else {
			count--
		}
	}
	return num
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
