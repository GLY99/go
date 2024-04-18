package main

import (
	"fmt"
)

func practice1() {
	var arr [26]byte
	var i byte = 0
	for ; i < 26; i++ {
		arr[i] = 'A' + i
		fmt.Printf("%c ", arr[i])
		if i == 25 {
			fmt.Printf("\n")
		}
	}
}

func practice2(arr *[5]int) int {
	var max int = (*arr)[0]
	for idx, val := range *arr {
		if idx == 0 {
			continue
		} else {
			if val > max {
				max = val
			}
		}
	}
	return max
}

func practice3(arr *[5]int) (sum int, avg float64) {
	for _, val := range *arr {
		sum += val
	}
	avg = float64(sum) / float64(len((*arr)))
	return
}

func main() {
	practice1()
	arr := [5]int{1, 5, 2, 4, 3}
	fmt.Printf("maxVal=%d\n", practice2(&arr))
	sum, avg := practice3(&arr)
	fmt.Printf("%d, %.2f\n", sum, avg)
}
