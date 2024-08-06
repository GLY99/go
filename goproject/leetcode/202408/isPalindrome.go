package main

func isPalindrome1(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	list := make([]int, 0)
	cur := head
	for cur != nil {
		list = append(list, cur.Val)
		cur = cur.Next
	}
	left := 0
	right := len(list) - 1
	for left < right {
		if list[left] != list[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isPalindrome(head *ListNode) bool {
	// 快慢指针找中点
	midNode := getMidNode(head)
	// 反转后半段链表
	newHead := reserveLink(midNode)
	c1 := head
	c2 := newHead
	// 奇数节点和偶数节点处理不同，所以这里再前半段遍历到最后一个节点时退出
	for c1 != nil && c1.Next != nil {
		if c1.Val != c2.Val {
			return false
		}
		c1 = c1.Next
		c2 = c2.Next
	}
	// 特殊判断奇数和偶数节点中间节点是否一样
	if c1 == c2 {
		return true
	}
	return c1.Val == c2.Val
}

func getMidNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reserveLink(head *ListNode) *ListNode {
	var preNode *ListNode
	preNode = nil
	cur := head
	for cur != nil {
		nextNode := cur.Next
		cur.Next = preNode
		preNode = cur
		cur = nextNode
	}
	return preNode
}
