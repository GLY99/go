package main

import (
	"fmt"
	demo "golangStudy/funcDemo/demos" // demo是别名
)

func main() {
	res := demo.Cal(1, 2, '+')
	fmt.Printf("%.2f\n", res)
}
