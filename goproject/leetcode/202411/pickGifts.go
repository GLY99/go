package main

import (
	"container/heap"
	"math"
)

type qp []int

func (q qp) Len() int {
	return len(q)
}

func (q qp) Less(i, j int) bool {
	return q[i] > q[j]
}

func (q qp) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *qp) Push(v interface{}) {
	*q = append(*q, v.(int))
}

func (q *qp) Pop() interface{} {
	n := len(*q)
	res := (*q)[n-1]
	*q = (*q)[:n-1]
	return res
}

func pickGifts(gifts []int, k int) int64 {
	q := qp{}
	for _, gift := range gifts {
		q.Push(gift)
	}
	heap.Init(&q)
	for k > 0 {
		x := heap.Pop(&q).(int)
		heap.Push(&q, int(math.Floor(math.Sqrt(float64(x)))))
		k--
	}
	var res int64 = 0
	for q.Len() > 0 {
		res += int64(q.Pop().(int))
	}
	return res
}
