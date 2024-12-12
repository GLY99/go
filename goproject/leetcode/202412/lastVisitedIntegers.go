package main

// https://leetcode.cn/problems/last-visited-integers/description/?envType=problem-list-v2&envId=simulation&status=TO_DO&difficulty=EASY
func lastVisitedIntegers(nums []int) []int {
	kNum := 0
	ans := make([]int, 0)
	seen := make([]int, 0)
	for _, num := range nums {
		if num != -1 {
			seen = append(seen, num)
			kNum = 0
		} else {
			kNum++
			if kNum <= len(seen) {
				ans = append(ans, seen[len(seen)-kNum])
			} else {
				ans = append(ans, -1)
			}
		}
	}
	return ans
}
