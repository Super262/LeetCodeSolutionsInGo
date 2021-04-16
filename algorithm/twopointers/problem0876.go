package twopointers

func middleNode(head *ListNode) *ListNode {
	pSlow := head
	pFast := head
	for pFast != nil && pFast.Next != nil {
		pSlow = pSlow.Next
		pFast = pFast.Next.Next
	}
	return pSlow
}
