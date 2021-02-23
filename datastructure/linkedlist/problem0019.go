package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}
	prev := (*ListNode)(nil)
	target := head
	tail := head
	for tail != nil && n > 0 {
		tail = tail.Next
		n--
	}
	if n > 0 {
		return head
	}
	for tail != nil {
		prev = target
		target = target.Next
		tail = tail.Next
	}
	if prev != nil {
		prev.Next = target.Next
	} else {
		head = head.Next
	}
	return head
}
