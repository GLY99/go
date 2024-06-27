package main

func findNumbers(nums []int) int {
	ans := 0
	for _, num := range nums {
		count := 0
		for num > 0 {
			num = num / 10
			count++
		}
		if count%2 == 0 {
			ans++
		}
	}
	return ans
}
