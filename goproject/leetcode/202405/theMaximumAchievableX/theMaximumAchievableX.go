package main

import "fmt"

func theMaximumAchievableX(num int, t int) int {
	return num + 2*t
}

func main() {
	fmt.Println(theMaximumAchievableX(1, 0))
}
