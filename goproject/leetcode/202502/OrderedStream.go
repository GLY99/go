package main

// https://leetcode.cn/problems/design-an-ordered-stream/?envType=daily-question&envId=2025-02-24

type OrderedStream struct {
	stream []string
	ptr    int
}

func Constructor3(n int) OrderedStream {
	return OrderedStream{stream: make([]string, n+1), ptr: 1}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.stream[idKey] = value
	start := this.ptr
	for this.ptr < len(this.stream) && this.stream[this.ptr] != "" {
		this.ptr++
	}
	return this.stream[start:this.ptr]
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
