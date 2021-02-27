package linkedlist

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p1 := head
	p2 := head
	for p2.Next != nil {
		p2 = p2.Next
		if p1.Val != p2.Val {
			p1.Next = p2
			p1 = p2
		}
	}
	p1.Next = nil
	return head
}
