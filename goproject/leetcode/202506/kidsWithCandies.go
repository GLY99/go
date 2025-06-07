package main

// https://leetcode.cn/problems/kids-with-the-greatest-number-of-candies/?envType=problem-list-v2&envId=array

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := candies[0]
	for _, candy := range candies {
		if candy > maxCandies {
			maxCandies = candy
		}
	}
	res := make([]bool, len(candies))
	for idx, candy := range candies {
		res[idx] = candy+extraCandies >= maxCandies
	}
	return res
}
