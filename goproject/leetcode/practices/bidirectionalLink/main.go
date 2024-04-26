package main

import "fmt"

// 带头的双向链表
type Link struct {
	Pre  *Link
	Next *Link
	Val  int
}

func insert(head *Link, node *Link) {
	curNode := head
	for curNode != nil {
		if curNode.Next == nil {
			curNode.Next = node
			node.Pre = curNode
			return
		} else if curNode.Val <= node.Val && curNode.Next.Val >= node.Val {
			node.Next = curNode.Next
			curNode.Next.Pre = node
			curNode.Next = node
			node.Pre = curNode
			return
		}
		curNode = curNode.Next
	}
}

func list(head *Link) {
	curNode := head.Next
	for curNode != nil {
		fmt.Println(curNode)
		curNode = curNode.Next
	}
}

func delete(head *Link, val int) {
	curNode := head.Next
	for curNode != nil {
		if curNode.Val == val {
			curNode.Pre.Next = curNode.Next
			curNode.Next = nil
		}
		curNode = curNode.Next
	}
}

func main() {
	head := Link{}
	node1 := Link{Val: 2}
	node2 := Link{Val: 1}
	node3 := Link{Val: 3}
	insert(&head, &node1)
	insert(&head, &node2)
	insert(&head, &node3)
	list(&head)
	delete(&head, 1)
	list(&head)
}
