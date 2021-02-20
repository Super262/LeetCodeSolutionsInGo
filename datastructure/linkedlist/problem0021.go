package linkedlist

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	queue := make([]*ListNode, 0)
	p1 := l1
	p2 := l2
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			queue = append(queue, p2)
			p2 = p2.Next
		} else {
			queue = append(queue, p1)
			p1 = p1.Next
		}
	}
	for p1 != nil {
		queue = append(queue, p1)
		p1 = p1.Next
	}
	for p2 != nil {
		queue = append(queue, p2)
		p2 = p2.Next
	}
	head := queue[0]
	newP := head
	for i := 1; i < len(queue); i++ {
		newP.Next = queue[i]
		newP = newP.Next
	}
	return head
}
