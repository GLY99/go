package main

import "sort"

// https://leetcode.cn/problems/WqXACV/?envType=problem-list-v2&envId=sorting

func halfQuestions(questions []int) int {
	quesMap := make([]int, 1001)
	for _, ques := range questions {
		quesMap[ques]++
	}
	sort.Ints(quesMap)
	var ans = 0
	n := len(questions) / 2
	i := 1000
	for n > 0 {
		ans++
		n -= quesMap[i]
		i--
	}
	return ans
}
