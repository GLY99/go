package main

import (
	"fmt"
	demo "golangStudy/funcDemo/demos" // demo是别名
)

func main() {
	res := demo.Cal(1, 2, '+')
	fmt.Printf("%.2f\n", res)
	sum, _ := demo.GetSumAndSub(1.2, 0.2)
	_, sub := demo.GetSumAndSub(1.2, 0.2)
	fmt.Printf("sum=%.2f, sub=%.2f", sum, sub)
}
