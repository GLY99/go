package main

import "fmt"

func bubbleSort(arr *[10]int) {
	// 从小到大冒泡排序
	length := len(*arr)
	for i := 0; i < length-1; i++ {
		flag := false
		for j := 0; j < length-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				tmp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func main() {
	var arr = [10]int{8, 2, 5, 6, 7, 1, 9, 3, 4, 0}
	fmt.Println(arr)
	bubbleSort(&arr)
	fmt.Println(arr)
}
