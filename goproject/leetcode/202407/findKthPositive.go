package main

func findKthPositive(arr []int, k int) int {
	idx := 0
	count := 0
	length := len(arr)
	for i := 1; i <= arr[length-1]; i++ {
		if i == arr[idx] {
			idx++
		} else {
			count++
			if count == k {
				return i
			}
		}
	}
	return arr[length-1] + k - count
}
