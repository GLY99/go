package main

import "fmt"

func exchangeArr(arr *[][]int) {
	for i := 0; i < len(*arr); i++ {
		for j := i + 1; j < len((*arr)[i]); j++ {
			tmpVal := (*arr)[i][j]
			(*arr)[i][j] = (*arr)[j][i]
			(*arr)[j][i] = tmpVal
		}
	}
}

func main() {
	// 1 2 3    1 4 7
	// 4 5 6 => 2 5 8
	// 7 8 9    3 6 9
	var testArr [][]int = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	exchangeArr(&testArr)
	fmt.Printf("%v\n", testArr)
}
