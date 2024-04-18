package main

import "fmt"

func func1(n int) (slice []int) {
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			slice = append(slice, 1)
		} else {
			slice = append(slice, slice[i-1]+slice[i-2])
		}
	}
	return
}

func main() {
	var n int = 10
	slice := func1(n)
	fmt.Println(slice)
}
