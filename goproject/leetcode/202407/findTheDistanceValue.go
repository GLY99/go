package main

import "sort"

func abs(num1, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	} else {
		return num2 - num1
	}
}

func findTheDistanceValue1(arr1 []int, arr2 []int, d int) int {
	count := 0
	for _, num1 := range arr1 {
		ok := true
		for _, num2 := range arr2 {
			if abs(num1, num2) <= d {
				ok = false
				break
			}
		}
		if ok {
			count++
		}
	}
	return count
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	// 先对arr2进行排序，之后遍历arr1，假设这个数为x，因为 d>=0
	// 如果arr2中存在一个数 y 且满足 num−d≤y≤num+d，那么arr1中的数x就不满足题目要求。
	sort.Ints(arr2)
	count := 0
	for _, num := range arr1 {
		low := num - d
		high := num + d
		if !binarySearch(arr2, low, high) {
			count++
		}
	}
	return count
}

func binarySearch(arr []int, low, high int) bool {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] >= low && arr[mid] <= high {
			return true
		} else if arr[mid] < low {
			left = mid + 1
		} else if arr[mid] > high {
			right = mid - 1
		}
	}
	return false
}
