package twopointers

func hasCycle(head *ListNode) bool {
	// 相当于追击问题
	if head == nil || head.Next == nil {
		return false
	}
	slowP := head
	fastP := head.Next
	for slowP != fastP {
		if fastP == nil || fastP.Next == nil {
			return false
		}
		slowP = slowP.Next
		fastP = fastP.Next.Next
	}
	return true
}
