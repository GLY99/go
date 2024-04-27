package main

import "fmt"

type HeroType struct {
	No       int
	Name     string
	Nickname string
	Next     *HeroType
}

// 给链表添加节点
func insert(head *HeroType, node *HeroType) {
	curNode := head
	for curNode.Next != nil {
		curNode = curNode.Next
	}
	curNode.Next = node
}

// 链表有序插入
func insertBySort(head *HeroType, node *HeroType) {
	curNode := head
	for curNode.Next != nil {
		if curNode.No == node.No {
			return
		} else if curNode.No < node.No && curNode.Next.No > node.No {
			node.Next = curNode.Next
			curNode.Next = node
			return
		} else {
			curNode = curNode.Next
		}
	}
	curNode.Next = node
}

func delete(head *HeroType, no int) {
	if head.Next == nil {
		return
	}
	curNode := head
	for curNode != nil {
		if curNode.Next != nil && curNode.Next.No == no {
			curNode.Next = curNode.Next.Next
			return
		} else {
			curNode = curNode.Next
		}
	}
}

// 遍历链表
func list(head *HeroType) {
	if head.Next == nil {
		return
	}
	curNode := head.Next
	for curNode != nil {
		fmt.Println(curNode)
		curNode = curNode.Next
	}
}

func main() {
	head := HeroType{}
	hero1 := HeroType{
		No: 2, Name: "songjiang", Nickname: "jishiyu", Next: nil,
	}
	insert(&head, &hero1)
	hero2 := HeroType{
		No: 3, Name: "lujunyi", Nickname: "laoliu", Next: nil,
	}
	insert(&head, &hero2)
	hero3 := HeroType{
		No: 1, Name: "me", Nickname: "me", Next: nil,
	}
	insertBySort(&head, &hero3)
	delete(&head, 1)
	list(&head)
}
