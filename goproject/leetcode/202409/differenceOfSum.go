package main

func abs(num1, num2 int) int {
	if num1 >= num2 {
		return num1 - num2
	}
	return num2 - num1
}

func differenceOfSum(nums []int) int {
	elementSum := 0
	numSum := 0
	for _, num := range nums {
		numSum += num
		for num > 0 {
			elementSum += num % 10
			num = num / 10
		}
	}
	return abs(elementSum, numSum)
}
