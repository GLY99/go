package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func insert(head *Node, node *Node) {
	if head.Next == nil {
		head.Val = node.Val
		head.Next = head
		return
	} else {
		curNode := head
		for curNode.Next != head {
			curNode = curNode.Next
		}
		node.Next = curNode.Next
		curNode.Next = node
	}
}

func delete(head *Node, val int) *Node {
	curNode := head
	helperNode := head
	if curNode.Next == nil {
		return head
	}
	if curNode.Next == head && curNode.Val == val {
		return &Node{}
	}
	for helperNode.Next != head {
		helperNode = helperNode.Next
	}
	for curNode.Next != head {
		if curNode.Val == val {
			helperNode.Next = curNode.Next
			if curNode == head {
				head = helperNode.Next
				return head
			} else {
				return helperNode.Next
			}
		}
		curNode = curNode.Next
		helperNode = helperNode.Next
	}
	if curNode.Val == val {
		helperNode.Next = curNode.Next
	}
	return head
}

func list(head *Node) {
	if head.Next == nil {
		return
	}
	curNode := head
	for curNode.Next != head {
		fmt.Printf("%p, %v\n", curNode, curNode)
		curNode = curNode.Next
	}
	fmt.Printf("%p, %v\n", curNode, curNode)
}

func main() {
	head := Node{}
	node1 := Node{Val: 1}
	node2 := Node{Val: 2}
	node3 := Node{Val: 3}
	insert(&head, &node1)
	insert(&head, &node2)
	insert(&head, &node3)
	list(&head)
	newhead := delete(&head, 1)
	list(newhead)
	newhead = delete(newhead, 2)
	list(newhead)
	newhead = delete(newhead, 3)
	list(newhead)
}
