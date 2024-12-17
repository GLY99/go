package main

func countEven(num int) int {
	var ans int
	for i := 1; i <= num; i++ {
		sum := 0
		for x := i; x > 0; x /= 10 {
			sum += x % 10
		}
		if sum%2 == 0 {
			ans++
		}
	}
	return ans
}
