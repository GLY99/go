package main

import (
	"fmt"
	"math"
)

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	res := []int{-1, -1}
label1:
	for i, num1 := range nums {
		for j, num2 := range nums {
			if int(math.Abs(float64(i-j))) >= indexDifference && int(math.Abs(float64(num1-num2))) >= valueDifference {
				res[0] = i
				res[1] = j
				break label1
			}
			if int(math.Abs(float64(j-i))) >= indexDifference && int(math.Abs(float64(num2-num1))) >= valueDifference {
				res[0] = j
				res[1] = i
				break label1
			}
		}
	}
	return res
}

func main() {
	fmt.Println(findIndices([]int{1, 2}, 1, 1))
}
