package main

import "fmt"

func insterNum(arr *[9]int, num int) [10]int {
	lengthArr := len(*arr)
	left := 0
	right := lengthArr - 1
	for {
		if left > right {
			break
		}
		mid := (left + right) / 2
		if (*arr)[mid] < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	insterIdx := left

	var newArr [10]int
	for i := 0; i < len(newArr); i++ {
		if i == insterIdx {
			newArr[i] = num
		} else if i < insterIdx {
			newArr[i] = (*arr)[i]
		} else {
			newArr[i] = (*arr)[i-1]
		}
	}
	return newArr
}

func main() {
	arr := [9]int{0, 1, 3, 4, 5, 6, 7, 8, 9}
	newArr := insterNum(&arr, 1)
	fmt.Println(newArr)
}
