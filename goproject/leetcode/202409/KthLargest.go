package main

import (
	"container/heap"
	"sort"
)

// 嵌套[]int类型
type KthLargest struct {
	k             int
	sort.IntSlice // 插入元素后自动从小到大排序
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, val := range nums {
		kl.Add(val)
	}
	return kl
}

func (this *KthLargest) Push(v interface{}) {
	this.IntSlice = append(this.IntSlice, v.(int))
}

func (this *KthLargest) Pop() interface{} {
	sl := this.IntSlice
	v := sl[len(sl)-1]
	this.IntSlice = sl[:len(sl)-1]
	return v
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}
	return this.IntSlice[0]
}
