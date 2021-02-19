package linkedlist

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prevNode := (*ListNode)(nil)
	currentNode := head
	nextNode := currentNode.Next
	for nextNode != nil {
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = nextNode
		nextNode = currentNode.Next
	}
	currentNode.Next = prevNode
	return currentNode
}
