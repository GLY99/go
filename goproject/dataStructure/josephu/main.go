package main

import "fmt"

type Person struct {
	Index int
	Next  *Person
}

// 构成单向环形链表
func makeLoopLine(num int) *Person {
	head := &Person{}
	tmpNode := &Person{}
	for i := 1; i <= num; i++ {
		node := &Person{Index: i}
		if i == 1 {
			head = node
		} else {
			tmpNode.Next = node
		}
		tmpNode = node
		if i == num {
			node.Next = head
		}
	}
	return head
}

func list(head *Person) int {
	if head.Next == nil {
		return 0
	}
	cur := head
	count := 0
	for cur.Next != head {
		fmt.Printf("%p, %v\n", cur, cur)
		cur = cur.Next
		count++
	}
	count++
	fmt.Printf("%p, %v\n", cur, cur)
	return count
}

func josephuGame(head *Person, startIdx int, step int) *Person {
	if head.Next == nil {
		return nil
	}
	if startIdx > list(head) {
		return nil
	}
	curNode := head
	PreNode := head
	for i := 1; i < startIdx; i++ {
		curNode = curNode.Next
	}
	for PreNode.Next != curNode {
		PreNode = PreNode.Next
	}
	for curNode.Next != curNode {
		for i := 1; i < step; i++ {
			curNode = curNode.Next
			PreNode = PreNode.Next
		}
		PreNode.Next = curNode.Next
		curNode = curNode.Next
	}
	return curNode
}

func main() {
	head := makeLoopLine(100)
	list(head)
	head = josephuGame(head, 1, 3) // 从第1个孩子开始数，数3个数
	list(head)                     // 91是最后留下来的
}
