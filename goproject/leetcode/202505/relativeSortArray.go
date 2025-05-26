package main

import "sort"

// https://leetcode.cn/problems/0H97ZC/?envType=problem-list-v2&envId=sorting

func relativeSortArray(arr1 []int, arr2 []int) []int {
	tempMap := make(map[int]int, 0)
	for _, val := range arr1 {
		tempMap[val]++
	}
	ans := make([]int, 0)
	notInArr2 := make([]int, 0)
	for _, val := range arr2 {
		if count, ok := tempMap[val]; ok {
			for i := 0; i < count; i++ {
				ans = append(ans, val)
				tempMap[val]--
			}
		}
	}
	for val, count := range tempMap {
		if count > 0 {
			for i := 0; i < count; i++ {
				notInArr2 = append(notInArr2, val)
			}
		}
	}
	sort.Ints(notInArr2)
	ans = append(ans, notInArr2...)
	return ans
}

func relativeSortArray1(arr1 []int, arr2 []int) []int {
	rank := make(map[int]int, 0)
	for idx, val := range arr2 {
		rank[val] = idx
	}
	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		rankX, hasX := rank[x]
		rankY, hasY := rank[y]
		if hasX && hasY {
			return rankX < rankY
		}
		if hasX || hasY {
			return hasX
		}
		return x < y
	})
	return arr1
}
