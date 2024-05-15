package main

import "fmt"

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

func PreOrder(root *Node) []string {
	res := []string{}
	var fc func(*Node)
	fc = func(node *Node) {
		if node != nil {
			res = append(res, node.Val)
			fc(node.Left)
			fc(node.Right)
		}
	}
	fc(root)
	return res
}

func InfixOrder(root *Node) []string {
	res := []string{}
	var fc func(*Node)
	fc = func(node *Node) {
		if node != nil {
			fc(node.Left)
			res = append(res, node.Val)
			fc(node.Right)
		}
	}
	fc(root)
	return res
}

func PostOrder(root *Node) []string {
	res := []string{}
	var fc func(*Node)
	fc = func(node *Node) {
		if node != nil {
			fc(node.Left)
			fc(node.Right)
			res = append(res, node.Val)
		}
	}
	fc(root)
	return res
}

func main() {
	root := Node{Val: "A"}
	node1 := Node{Val: "B"}
	node2 := Node{Val: "C"}
	root.Left = &node1
	root.Right = &node2
	node3 := Node{Val: "D"}
	node4 := Node{Val: "E"}
	node5 := Node{Val: "F"}
	node6 := Node{Val: "G"}
	node1.Left = &node3
	node1.Right = &node4
	node2.Left = &node5
	node2.Right = &node6
	node7 := Node{Val: "H"}
	node3.Right = &node7
	fmt.Println(PreOrder(&root))
	fmt.Println(InfixOrder(&root))
	fmt.Println(PostOrder(&root))
}
