package main

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	set := make(map[int]interface{})
	tmpHead := &ListNode{}
	tmpHead.Next = head
	cur := head
	pre := tmpHead
	for cur != nil {
		_, ok := set[cur.Val]
		if ok {
			pre.Next = cur.Next
		} else {
			set[cur.Val] = true
			pre = cur
		}
		cur = cur.Next
	}
	return tmpHead.Next
}
