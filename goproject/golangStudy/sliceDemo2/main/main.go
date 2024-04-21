package main

import "fmt"

func main() {
	sliceA := []int{1, 2, 3}
	sliceB := sliceA[:2]
	sliceA[0] = 0
	fmt.Println(sliceA, sliceB) // [0 2 3] [0 2]
}
