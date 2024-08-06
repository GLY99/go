package main

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var preNode *ListNode
	preNode = nil
	curNode := head
	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = nextNode
	}
	return preNode
}
