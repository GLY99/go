package main

import "fmt"

func binarySearch(arr []int, n int) bool {
	var length = len(arr)
	var left = 0
	var right = length - 1
	for {
		if left > right {
			break
		}
		mid := (left + right) / 2
		if arr[mid] == n {
			return true
		} else if arr[mid] < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func main() {
	var arr1 = []int{0, 1, 2, 3, 4}
	for i := -1; i <= 5; i++ {
		fmt.Println(binarySearch(arr1, i))
	}

	arr1 = []int{0, 1, 2, 3, 4, 5}
	for i := -1; i <= 6; i++ {
		fmt.Println(binarySearch(arr1, i))
	}
}
