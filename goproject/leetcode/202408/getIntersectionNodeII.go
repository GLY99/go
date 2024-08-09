package main

func getIntersectionNodeII(headA, headB *ListNode) *ListNode {
	aLength := 0
	bLength := 0
	curA := headA
	curB := headB
	for curA != nil || curB != nil {
		if curA != nil {
			aLength++
			curA = curA.Next
		}
		if curB != nil {
			bLength++
			curB = curB.Next
		}
	}
	curA = headA
	curB = headB
	if aLength > bLength {
		tmp := aLength - bLength
		for tmp > 0 {
			curA = curA.Next
			tmp--
		}
	} else if aLength < bLength {
		tmp := bLength - aLength
		for tmp > 0 {
			curB = curB.Next
			tmp--
		}
	}
	for curA != nil && curB != nil {
		if curA == curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}
	return nil
}
