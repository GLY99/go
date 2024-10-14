package main

func arithmeticTriplets(nums []int, diff int) int {
	mapping := make(map[int]int)
	for _, num := range nums {
		mapping[num]++
	}
	ans := 0
	for _, num := range nums {
		_, ok1 := mapping[num+diff]
		_, ok2 := mapping[num+diff*2]
		if ok1 && ok2 {
			ans++
		}
	}
	return ans
}
