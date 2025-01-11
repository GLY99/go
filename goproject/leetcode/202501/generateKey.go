package main

func generateKey(num1 int, num2 int, num3 int) int {
	res := 0
	for i := 1; num1 > 0 && num2 > 0 && num3 > 0; i *= 10 {
		res += min(num1%10, min(num2%10, num3%10)) * i
		num1, num2, num3 = num1/10, num2/10, num3/10
	}
	return res
}
