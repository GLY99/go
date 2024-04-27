package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	front   int
	rear    int
	maxsize int
	arr     [5]int
}

func (queue *Queue) AddQueue(num int) (err error) {
	if queue.rear == queue.maxsize-1 {
		err = errors.New("queue full")
		return
	}
	queue.rear++
	queue.arr[queue.rear] = num
	return
}

func (queue *Queue) GetQueue() (num int, err error) {
	if queue.front == queue.rear {
		return -1, errors.New("queue empty")
	}
	queue.front++
	return queue.arr[queue.front], nil
}

func (queue *Queue) ListQueue() {
	for i := queue.front + 1; i <= queue.rear; i++ {
		fmt.Println(queue.arr[i])
	}
}

func main() {
	queue := &Queue{
		front: -1, rear: -1, maxsize: 5,
		arr: *new([5]int),
	}
	for i := 0; i < 6; i++ {
		err := queue.AddQueue(1)
		fmt.Println(err)
	}
	queue.ListQueue()
	for i := 0; i < 6; i++ {
		num, err := queue.GetQueue()
		fmt.Println(num, err)
	}
}
