package main

func countValidSelections(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	pre := 0
	ans := 0
	for _, num := range nums {
		if num > 0 {
			pre += num
		} else if pre*2 == total {
			ans += 2
		} else if pre*2-total == 1 || pre*2-total == -1 {
			ans++
		}
	}
	return ans
}
