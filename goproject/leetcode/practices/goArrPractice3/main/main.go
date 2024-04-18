package main

import (
	"fmt"
	"math/rand"
)

func getMaxAndMinVal() (min int, max int) {
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		num := rand.Intn(100)
		arr[i] = num
	}
	min = arr[0]
	max = arr[0]
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("%d,", arr[i])
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	fmt.Println()
	return
}

func main() {
	min, max := getMaxAndMinVal()
	fmt.Printf("%d, %d", min, max)
}
