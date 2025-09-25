package main

// https://leetcode.cn/problems/the-employee-that-worked-on-the-longest-task/?envType=problem-list-v2&envId=array

func hardestWorker(n int, logs [][]int) int {
	startTime := 0
	id := 0
	maxTime := 0
	for _, log := range logs {
		costTime := log[1] - startTime
		if costTime > maxTime || costTime == maxTime && log[0] < id {
			maxTime = costTime
			id = log[0]
		}
		startTime = log[1]
	}
	return id
}
