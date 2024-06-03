package main

func distributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people)
	idx := 0
	count := 1
	for candies > 0 {
		if count > candies {
			res[idx%num_people] += candies
			break
		}
		res[idx%num_people] += count
		candies -= count
		count++
		idx++
	}
	return res
}
