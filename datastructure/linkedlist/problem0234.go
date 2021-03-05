package linkedlist

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	listLen := 0
	p1 := head
	for p1 != nil {
		listLen++
		p1 = p1.Next
	}
	mid := listLen / 2
	p0 := (*ListNode)(nil)
	p1 = head
	p2 := head.Next
	for i := 0; i < mid; i++ {
		p1.Next = p0
		p0 = p1
		p1 = p2
		p2 = p2.Next
	}
	if listLen%2 == 1 {
		p1 = p1.Next
	}
	for p0 != nil && p1 != nil {
		if p0.Val == p1.Val {
			p0 = p0.Next
			p1 = p1.Next
		} else {
			break
		}
	}
	return p0 == nil && p1 == nil
}
