package doublepointers

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slowP := head.Next
	if slowP == nil {
		return false
	}
	if slowP.Next == nil {
		return false
	}
	fastP := slowP.Next
	for slowP != nil && fastP != nil {
		if fastP == slowP {
			return true
		} else {
			slowP = slowP.Next
			fastP = fastP.Next
			if fastP == nil {
				return false
			} else {
				fastP = fastP.Next
			}
		}

	}
	return false
}
