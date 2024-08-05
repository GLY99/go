package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	aLength := 0
	bLength := 0
	curANode := headA
	curBNode := headB
	for curANode != nil || curBNode != nil {
		if curANode != nil {
			aLength++
			curANode = curANode.Next
		}
		if curBNode != nil {
			bLength++
			curBNode = curBNode.Next
		}
	}
	curANode = headA
	curBNode = headB
	x := abs(aLength, bLength)
	if x != 0 && aLength > bLength {
		for i := 0; i < x; i++ {
			curANode = curANode.Next
		}
	}
	if x != 0 && aLength < bLength {
		for i := 0; i < x; i++ {
			curBNode = curBNode.Next
		}
	}
	for curANode != nil && curBNode != nil {
		if curANode == curBNode {
			return curANode
		}
		curANode = curANode.Next
		curBNode = curBNode.Next
	}
	return nil
}

func abs(num1, num2 int) int {
	if num1 >= num2 {
		return num1 - num2
	}
	return num2 - num1
}
