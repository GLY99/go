package main

import "fmt"

func updateArr(arr *[3][4]int, n int) {
	rows := len((*arr))
	cols := len((*arr)[0])
	if n > rows*cols {
		return
	}
	var rowIdx int
	var colIdx int
	if n%cols == 0 {
		rowIdx = n/cols - 1
		colIdx = cols - 1
	} else {
		rowIdx = n / cols
		colIdx = (n % cols) - 1
	}
	if rowIdx-1 >= 0 {
		(*arr)[rowIdx-1][colIdx] = 0
		if colIdx-1 >= 0 {
			(*arr)[rowIdx-1][colIdx-1] = 0
			(*arr)[rowIdx][colIdx-1] = 0
		}
		if colIdx+1 < cols {
			(*arr)[rowIdx-1][colIdx+1] = 0
			(*arr)[rowIdx][colIdx+1] = 0
		}
	}
	if rowIdx+1 < rows {
		(*arr)[rowIdx+1][colIdx] = 0
		if colIdx-1 >= 0 {
			(*arr)[rowIdx+1][colIdx-1] = 0
		}
		if colIdx+1 < cols {
			(*arr)[rowIdx+1][colIdx+1] = 0
		}
	}
}

func main() {
	var arr = [3][4]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}
	updateArr(&arr, 8)
	fmt.Println(arr)
}
