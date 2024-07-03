package main

func replaceElements(arr []int) []int {
	length := len(arr)
	curMax := arr[length-1]
	arr[length-1] = -1
	for i := length - 2; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = curMax
		curMax = max(curMax, tmp)
	}
	return arr
}
