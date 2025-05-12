package main

// https://leetcode.cn/problems/finding-3-digit-even-numbers/?envType=daily-question&envId=2025-05-12

func findEvenNumbers(digits []int) []int {
	mapping := make(map[int]int)
	for _, digit := range digits {
		mapping[digit]++
	}
	ans := []int{}
	for i := 100; i <= 998; i += 2 {
		tempMapping := map[int]int{}
		temp := i
		for temp > 0 {
			tempMapping[temp%10]++
			temp /= 10
		}
		flag := true
		for k, v := range tempMapping {
			if mapping[k] < v {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, i)
		}
	}
	return ans
}
