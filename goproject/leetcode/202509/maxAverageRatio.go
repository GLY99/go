package main

import "container/heap"

// https://leetcode.cn/problems/maximum-average-pass-ratio/?envType=daily-question&envId=2025-09-01

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := hp(classes)
	heap.Init(h)
	for ; extraStudents > 0; extraStudents-- {
		h[0][0]++
		h[0][1]++
		heap.Fix(&h, 0)
	}
	var ans float64 = 0
	for _, c := range h {
		ans += float64(c[0]) / float64(c[1])
	}
	return ans / float64(len(classes))
}

type hp [][]int

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	// (totali + 1)×(totali)×(totalj − passj) < (totalj + 1)×(totalj)×(totali − passi) 时
	// idxi的优先级 > idxj; 放到函数内部就是idxj的优先级大于idxj-1交换idxj和idxj-1的位置
	// func insertionSort(data Interface, a, b int) {
	//     for i := a + 1; i < b; i++ {
	//         for j := i; j > a && data.Less(j, j-1); j-- {
	//             data.Swap(j, j-1)
	//         }
	//     }
	// }
	a := h[i]
	b := h[j]
	return (a[1]+1)*a[1]*(b[1]-b[0]) < (b[1]+1)*b[1]*(a[1]-a[0])
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }
