package main

import "sort"

type s struct {
	line  int
	count int
}

type sSlice []*s

func (a sSlice) Len() int {
	return len(a)
}
func (a sSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a sSlice) Less(i, j int) bool {
	if a[i].count != a[j].count {
		return a[i].count < a[j].count
	}
	return a[i].line < a[j].line
}

func kWeakestRows(mat [][]int, k int) []int {
	arr := make(sSlice, 0)
	for i := 0; i < len(mat); i++ {
		count := 0
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				count++
			}
		}
		arr = append(arr, &s{i, count})
	}
	sort.Sort(arr)
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = arr[i].line
	}
	return res
}
