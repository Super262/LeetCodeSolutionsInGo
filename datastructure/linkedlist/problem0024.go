package linkedlist

func swapPairs(head *ListNode) *ListNode {
	p1 := head
	var p2 *ListNode
	var t int
	for p1 != nil {
		p2 = p1.Next
		if p2 == nil {
			return head
		}
		t = p1.Val
		p1.Val = p2.Val
		p2.Val = t
		p1 = p2.Next
	}
	return head
}
