package main

func addDigits(num int) int {
	for num > 9 {
		tmpNum := 0
		for num != 0 {
			tmpNum += num % 10
			num = num / 10
		}
		num = tmpNum
	}
	return num
}
