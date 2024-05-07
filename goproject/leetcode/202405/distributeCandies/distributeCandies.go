package main

// https://leetcode.cn/problems/distribute-candies/

func distributeCandies(candyType []int) int {
	set := make(map[int]struct{})
	for _, t := range candyType {
		set[t] = struct{}{}
	}
	ans := len(candyType) / 2
	if len(set) < ans {
		ans = len(set)
	}
	return ans
}

func main() {
	distributeCandies([]int{1, 1, 2, 3, 3})
}
