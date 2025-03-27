package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 || head == nil {
		return head
	}
	fast := head
	slow := head
	for n > 0 {
		if fast != nil {
			fast = fast.Next
			n--
		} else {
			return head
		}
	}
	if fast == nil {
		return head.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// slow是第n-1个结点
	slow.Next = slow.Next.Next
	return head
}
