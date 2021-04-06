package hashtable

func copyRandomList(head *Node) *Node {
	oldToNew := make(map[*Node]*Node)
	currentNode := head
	var prevCopied *Node
	for currentNode != nil {
		currentCopied := &Node{Val: currentNode.Val}
		oldToNew[currentNode] = currentCopied
		if prevCopied != nil {
			prevCopied.Next = currentCopied
		}
		currentNode = currentNode.Next
		prevCopied = currentCopied
	}
	currentNode = head
	for currentNode != nil {
		if currentNode.Random != nil {
			oldToNew[currentNode].Random = oldToNew[currentNode.Random]
		}
		currentNode = currentNode.Next
	}
	return oldToNew[head]
}
