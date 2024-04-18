package main

import "fmt"

func exchangeArr(arr *[4][4]int) {
	rows := len(*arr)
	cols := len((*arr)[0])
	for i := 0; i <= rows/2; i++ {
		for j := 0; j < cols; j++ {
			tmp := (*arr)[i][j]
			(*arr)[i][j] = (*arr)[rows-i-1][j]
			(*arr)[rows-i-1][j] = tmp
		}
	}
}

func main() {
	var arr = [4][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16}}
	exchangeArr(&arr)
	fmt.Println(arr)
}
