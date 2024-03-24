package main

import "fmt"

type Circle struct {
	radius float64
}

func (circle *Circle) area() float64 {
	// 编译器底层做了优化也可以写成circle.radius
	return 3.14 * (*circle).radius * (*circle).radius
}

func main() {
	var circle Circle
	circle.radius = 1.0
	// 编译器底层做了优化也可以写成circle.area()
	area := (&circle).area()
	fmt.Println(area)
}
