package main

import "fmt"

func swapNumbers(numbers []int) []int {
	numbers[0] = numbers[0] + numbers[1] // {3, 2}
	numbers[1] = numbers[0] - numbers[1] // {3, 1}
	numbers[0] -= numbers[1]
	return numbers
}

func main() {
	numbers := []int{1, 2}
	fmt.Println(swapNumbers(numbers))
}
