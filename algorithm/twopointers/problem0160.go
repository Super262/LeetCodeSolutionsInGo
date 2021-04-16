package twopointers

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1 := headA
	p2 := headB
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
		if p1 == nil && p2 == nil {
			return nil
		}
		if p1 == nil {
			p1 = headB
		}
		if p2 == nil {
			p2 = headA
		}
	}
	return p1
}
