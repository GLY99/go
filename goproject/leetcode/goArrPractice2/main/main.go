package main

import (
	"fmt"
	"math/rand"
)

func practice1(arr *[5]int) {
	arr_length := len((*arr))
	for i := 0; i < arr_length/2; i++ {
		tmp := arr[i]
		arr[i] = arr[arr_length-1-i]
		arr[arr_length-1-i] = tmp
	}
}

func main() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Printf("%v\n", arr)
	practice1(&arr)
	fmt.Printf("%v\n", arr)
}
