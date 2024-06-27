package main

func findSpecialInteger(arr []int) int {
	cur, count := arr[0], 0
	length := len(arr)
	for i := 0; i < length; i++ {
		if arr[i] == cur {
			count++
			if count*4 > length {
				return cur
			}
		} else {
			cur, count = arr[i], 1
		}
	}
	return -1
}
