package main

func kthToLast1(head *ListNode, k int) int {
	length := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		length++
	}
	length = length - k
	cur = head
	for length > 0 {
		cur = cur.Next
		length--
	}
	return cur.Val
}

func kthToLast(head *ListNode, k int) int {
	fast := head
	slow := head
	for k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}
