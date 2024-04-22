package main

import (
	"errors"
	"fmt"
)

type LoopQueue struct {
	front   int
	rear    int
	maxsize int
	arr     [5]int
}

func (queue *LoopQueue) AddQueue(num int) (err error) {
	if (queue.rear+1)%queue.maxsize == queue.front {
		err = errors.New("queue full")
		return
	}
	queue.arr[queue.rear] = num
	queue.rear = (queue.rear + 1) % queue.maxsize
	return
}

func (queue *LoopQueue) GetQueue() (num int, err error) {
	if queue.front == queue.rear {
		return -1, errors.New("queue empty")
	}
	num = queue.arr[queue.front]
	queue.front = (queue.front + 1) % queue.maxsize
	return num, nil
}

func (queue *LoopQueue) ListQueue() {
	size := queue.QueueSize()
	if size == 0 {
		fmt.Printf("队列是空的\n")
	}
	for i := queue.front; i < 2*(queue.maxsize-1); i++ {
		fmt.Println(queue.arr[i%queue.maxsize])
		if i%queue.maxsize == queue.rear-1 {
			break
		}
	}
}

func (queue *LoopQueue) QueueSize() int {
	return (queue.rear + queue.maxsize - queue.front) % queue.maxsize
}

func main() {
	queue := &LoopQueue{
		front: 0, rear: 0, maxsize: 5,
		arr: *new([5]int),
	}
	for i := 0; i < 5; i++ {
		err := queue.AddQueue(1)
		fmt.Println(err)
	}
	queue.ListQueue()
	for i := 0; i < 5; i++ {
		num, err := queue.GetQueue()
		fmt.Println(num, err)
	}
	for i := 0; i < 5; i++ {
		err := queue.AddQueue(1)
		fmt.Println(err)
	}
	queue.ListQueue()
	fmt.Println(queue.QueueSize())
}
