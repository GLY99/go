package main

import "sort"

func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	length := len(arr)
	for i := 0; i < length; i++ {
		j := i + 1
		for j < length {
			if arr[j] == 2*arr[i] {
				return true
			} else if arr[j] > 2*arr[i] {
				break
			}
			j++
		}
	}
	for i := length - 1; i >= 0; i-- {
		j := i - 1
		for j >= 0 {
			if arr[j] == 2*arr[i] {
				return true
			} else if arr[j] < 2*arr[i] {
				break
			}
			j--
		}
	}
	return false
}
