package main

type TripleInOne struct {
	size   int
	lens   *[3]int
	stacks []int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{size: stackSize, lens: new([3]int), stacks: make([]int, stackSize*3)}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	if this.lens[stackNum] == this.size {
		return
	}
	this.stacks[this.size*stackNum+(*this.lens)[stackNum]] = value
	this.lens[stackNum]++
}

func (this *TripleInOne) Pop(stackNum int) int {
	if this.lens[stackNum] == 0 {
		return -1
	}
	val := this.stacks[this.size*stackNum+(*this.lens)[stackNum]-1]
	this.stacks[this.size*stackNum+(*this.lens)[stackNum]-1] = 0
	this.lens[stackNum]--
	return val
}

func (this *TripleInOne) Peek(stackNum int) int {
	if this.lens[stackNum] == 0 {
		return -1
	}
	val := this.stacks[this.size*stackNum+(*this.lens)[stackNum]-1]
	return val
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.lens[stackNum] == 0
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */
