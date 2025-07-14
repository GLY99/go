package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	curNode := head
	ans := 0
	for curNode != nil {
		ans = ans*2 + curNode.Val
		curNode = curNode.Next
	}
	return ans
}
