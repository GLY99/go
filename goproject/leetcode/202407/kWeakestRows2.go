package main

import (
	"container/heap"
	"sort"
)

type Pair struct {
	line  int
	count int
}

type Hp []*Pair

func (h Hp) Len() int {
	return len(h)
}

func (h Hp) Less(i, j int) bool {
	if h[i].count != h[j].count {
		return h[i].count < h[j].count
	}
	return h[i].line < h[j].line
}

func (h Hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Hp) Push(v interface{}) {
	*h = append(*h, v.(*Pair))
}

func (h *Hp) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func kWeakestRows2(mat [][]int, k int) []int {
	h := Hp{}
	for i, row := range mat {
		count := sort.Search(len(row), func(j int) bool { return row[j] == 0 })
		h = append(h, &Pair{i, count})
	}
	heap.Init(&h)
	ans := make([]int, k)
	for i := range ans {
		ans[i] = heap.Pop(&h).(*Pair).line
	}
	return ans
}
