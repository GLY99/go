package main

//https://leetcode.cn/problems/count-distinct-numbers-on-board/description/?envType=problem-list-v2&envId=simulation&status=TO_DO&difficulty=EASY

func distinctIntegers(n int) int {
	// nums := make([]int, n+1)
	// nums[n] = 1
	// for i := 0; i < n; i++ {
	// 	for j := 1; j < n; j++ {
	// 		if nums[j] == 0 {
	// 			continue
	// 		}
	// 		for k := 1; k < n; k++ {
	// 			if j%k == 1 {
	// 				nums[k] = 1
	// 			}
	// 		}
	// 	}
	// }
	// var res int
	// for _, num := range nums {
	// 	res += num
	// }
	// return res
	return max(n-1, 1)
}
